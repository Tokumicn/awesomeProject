package langchian_spliter

import (
	"fmt"
	"github.com/tmc/langchaingo/textsplitter"
	"testing"
)

func TestNewTokenSplitter(t *testing.T) {
	splitter := textsplitter.NewTokenSplitter(
		textsplitter.WithChunkSize(20),   // 每块约200字符
		textsplitter.WithChunkOverlap(2), // 块间重叠20字符
	)

	text := "这是第一篇段落。包含多个句子！第二篇段落...（长文本） 这是第二篇段落。包含多个句子！第二篇段落...（长文本） 这是第三篇段落。包含多个句子！第二篇段落...（长文本）"
	chunks, err := splitter.SplitText(text)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("len:%d, chunks: %+v", len(chunks), chunks))
}
