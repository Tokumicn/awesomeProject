package cr_spliter

import (
	"regexp"
	"strings"
)

type ChineseRecursiveSplitter struct {
	ChunkSize    int
	ChunkOverlap int
	Separators   []string
}

func NewChineseRecursiveSplitter() *ChineseRecursiveSplitter {
	return &ChineseRecursiveSplitter{
		ChunkSize:    400,
		ChunkOverlap: 20,
		Separators:   []string{"\n\n", "\n", "。", "！", "？", "，", " "},
	}
}

func (s *ChineseRecursiveSplitter) Split(text string) []string {
	// 预处理：合并多余空白符
	re := regexp.MustCompile(`\s+`)
	cleaned := re.ReplaceAllString(text, " ")

	// 递归分割逻辑
	return s.recursiveSplit(cleaned, s.Separators)
}

func (s *ChineseRecursiveSplitter) recursiveSplit(text string, separators []string) []string {
	if len(separators) == 0 || len(text) <= s.ChunkSize {
		return []string{text}
	}

	sep := separators[0]
	parts := strings.Split(text, sep)

	var result []string
	currentChunk := ""

	for _, part := range parts {
		if len(currentChunk)+len(part)+len(sep) > s.ChunkSize {
			if currentChunk != "" {
				result = append(result, currentChunk)
				currentChunk = currentChunk[len(currentChunk)-s.ChunkOverlap:] + sep + part
			} else {
				currentChunk = part
			}
		} else {
			currentChunk += sep + part
		}
	}

	if currentChunk != "" {
		result = append(result, currentChunk)
	}

	// 若分割后仍超限，尝试下一级分隔符
	if len(result) == 0 || len(result[0]) > s.ChunkSize {
		return s.recursiveSplit(text, separators[1:])
	}

	return result
}
