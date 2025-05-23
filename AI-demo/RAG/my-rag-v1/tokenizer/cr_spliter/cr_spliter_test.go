package cr_spliter

import (
	"fmt"
	"testing"
)

func TestChineseRecursiveSplitter_Split(t *testing.T) {
	splitter := NewChineseRecursiveSplitter()
	text := "预训练模型通过大规模数据学习通用特征...（长文本示例）"
	chunks := splitter.Split(text)
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d:\n%s\n\n", i+1, chunk)
	}
}
