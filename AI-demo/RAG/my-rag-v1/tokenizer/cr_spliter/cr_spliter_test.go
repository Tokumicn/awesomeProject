package cr_spliter

import (
	"fmt"
	"testing"
)

func TestChineseRecursiveSplitter_Split(t *testing.T) {
	splitter := NewChineseRecursiveSplitter()
	texts := []string{
		"鱼钩上不放鱼饵能钓到吗？为什么",
		"八卦是什么？",
	}

	for _, text := range texts {
		chunks := splitter.Split(text)
		for i, chunk := range chunks {
			fmt.Printf("Chunk %d:\n%s\n\n", i+1, chunk)
		}
	}
}
