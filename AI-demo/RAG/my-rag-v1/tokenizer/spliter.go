package tokenizer

import (
	"ai-demo/RAG/my-rag-v1/tokenizer/cr_spliter"
	"ai-demo/RAG/my-rag-v1/tokenizer/jieba_spliter"
	"ai-demo/RAG/my-rag-v1/tokenizer/langchian_spliter"
)

type Options struct {
	ChunkSize    int
	ChunkOverlap int
	SplitterType SplitterType
	CutType      string
}

// Splitter 定义分词器接口
type Splitter interface {
	Split(text string) []string
}

// cutType 分词类型
type SplitterType string

const (
	None      SplitterType = "None"
	Jieba     SplitterType = "JieBa"
	CR        SplitterType = "ChineseRecursive"
	LangChain SplitterType = "LangChain"
)

// 统一封装的分词器结构体
type MultiSplitter struct {
	op Options
}

// NewSplitter 创建统一分词器实例
func NewSplitter(op Options) *MultiSplitter {
	return &MultiSplitter{
		op: op,
	}
}

// Split 根据类型选择不同的分词实现
func (ms *MultiSplitter) Split(text string) []string {
	switch ms.op.SplitterType {
	case Jieba:
		return jieba_spliter.NewSplitter(true).Split(ms.op.CutType, text)
	case CR:
		return cr_spliter.NewChineseRecursiveSplitter().Split(text)
	case LangChain:
		splitter := langchian_spliter.NewChineseRecursiveSplitter(ms.op.ChunkSize, ms.op.ChunkOverlap)
		return splitter.Split(text)
	case None:
		return []string{text}
	default:
		panic("unsupported cut type")
	}
}
