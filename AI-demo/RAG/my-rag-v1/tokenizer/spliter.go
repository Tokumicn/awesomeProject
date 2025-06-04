package tokenizer

import (
	"ai-demo/RAG/my-rag-v1/tokenizer/cr_spliter"
	"ai-demo/RAG/my-rag-v1/tokenizer/jieba_spliter"
)

// Splitter 定义分词器接口
type Splitter interface {
	Split(text string) []string
}

// cutType 分词类型
type SplitterType string

const (
	None  SplitterType = "None"
	Jieba SplitterType = "JieBa"
	CR    SplitterType = "ChineseRecursive"
)

// 统一封装的分词器结构体
type MultiSplitter struct{}

// NewSplitter 创建统一分词器实例
func NewSplitter() *MultiSplitter {
	return &MultiSplitter{}
}

// Split 根据类型选择不同的分词实现
func (ms *MultiSplitter) Split(splitterType SplitterType, cutType string, text string) []string {
	switch splitterType {
	case Jieba:
		return jieba_spliter.NewSplitter(true).Split(cutType, text)
	case CR:
		return cr_spliter.NewChineseRecursiveSplitter().Split(text)
	case None:
		return []string{text}
	default:
		panic("unsupported cut type")
	}
}
