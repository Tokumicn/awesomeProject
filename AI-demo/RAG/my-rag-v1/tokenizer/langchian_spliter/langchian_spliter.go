package langchian_spliter

import (
	"github.com/tmc/langchaingo/textsplitter"
	"log/slog"
	"regexp"
)

type LangChainRecursiveSplitter struct {
	ChunkSize    int
	ChunkOverlap int
	splitter     textsplitter.TextSplitter
}

func NewChineseRecursiveSplitter(chunkSize, chunkOverlap int) *LangChainRecursiveSplitter {
	splitter := textsplitter.NewTokenSplitter(
		textsplitter.WithChunkSize(200),   // 每块约200字符
		textsplitter.WithChunkOverlap(20), // 块间重叠20字符
	)

	return &LangChainRecursiveSplitter{
		ChunkSize:    chunkSize,
		ChunkOverlap: chunkOverlap,
		splitter:     splitter,
	}
}

func (s *LangChainRecursiveSplitter) Split(text string) []string {
	// 预处理：合并多余空白符
	re := regexp.MustCompile(`\s+`)
	cleaned := re.ReplaceAllString(text, " ")

	strings, err := s.splitter.SplitText(cleaned)
	if err != nil {
		slog.Error("LangChainRecursiveSplitter split text failed", "error", err)
		return nil
	}
	return strings
}
