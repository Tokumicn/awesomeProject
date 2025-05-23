package jieba_spliter

import (
	"fmt"
	"log/slog"

	"github.com/wangbin/jiebago"
)

// Python原版实现：https://github.com/fxsjy/jieba
// 核心算法：
// 1. 基于前缀词典实现高效的词图扫描，生成句子中汉字所有可能成词情况所构成的有向无环图 (DAG)
// 2. 采用了动态规划查找最大概率路径, 找出基于词频的最大切分组合
// 3. 对于未登录词，采用了基于汉字成词能力的 HMM 模型，使用了 Viterbi 算法

// Go版本实现：https://github.com/wangbin/jiebago

const (
	// CutModeAccurate 精确模式，试图将句子最精确地切开，适合文本分析
	CutModeAccurate = "accurate"
	// CutModeAccurateNew 新词模式  依然是 Cut 仅 hhm=true
	CutModeAccurateNew = "accurate_new"
	// CutModeAll 全模式，把句子中所有的可以成词的词语都扫描出来, 速度非常快，但是不能解决歧义不准确
	CutModeAll = "all"
	// CutModeSearch 搜索引擎模式，在精确模式（accurate）的基础上，对长词再次切分，提高召回率，适合用于搜索引擎分词
	CutModeSearch = "search"
)

type JieBaSeg struct {
	segmenter *jiebago.Segmenter
	cutChan   <-chan string
	hmm       bool
}

func NewSplitter(hmm bool) JieBaSeg {
	jbSeg := JieBaSeg{
		segmenter: &jiebago.Segmenter{},
		cutChan:   make(<-chan string),
		hmm:       hmm,
	}

	err := jbSeg.segmenter.LoadDictionary("./tmp/dict.txt")
	if err != nil {
		slog.Error("jiebago load dictionary failed")
	}
	return jbSeg
}

func (seg JieBaSeg) PrintCh(ch <-chan string) {
	words := make([]string, 0)
	for word := range ch {
		words = append(words, word)
	}

	seg.Print(words)
}

func (seg JieBaSeg) Print(words []string) {
	fmt.Println(words)
}

func (seg JieBaSeg) Split(cutType, text string) []string {

	switch cutType {
	case CutModeAccurate:
		cutCh := seg.segmenter.Cut(text, false)
		seg.cutChan = cutCh
	case CutModeAccurateNew:
		cutCh := seg.segmenter.Cut(text, true)
		seg.cutChan = cutCh
	case CutModeAll:
		cutCh := seg.segmenter.CutAll(text)
		seg.cutChan = cutCh
	case CutModeSearch:
		cutCh := seg.segmenter.CutForSearch(text, seg.hmm)
		seg.cutChan = cutCh
	default:
		cutCh := seg.segmenter.Cut(text, seg.hmm)
		seg.cutChan = cutCh
	}

	var result []string
	for word := range seg.cutChan {
		result = append(result, word)
	}
	return result
}
