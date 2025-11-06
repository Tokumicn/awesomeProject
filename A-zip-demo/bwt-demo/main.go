package main

import (
	"fmt"
	"sort"
)

// BWT 正向变换
func BWT(input string) string {
	if len(input) == 0 {
		return ""
	}

	// 添加终止符
	s := input + "$"
	n := len(s)

	// 生成循环移位索引
	rotations := make([]int, n)
	for i := range rotations {
		rotations[i] = i
	}

	// 按字典序排序移位
	sort.Slice(rotations, func(i, j int) bool {
		a, b := rotations[i], rotations[j]
		for k := 0; k < n; k++ {
			ai, bj := (a+k)%n, (b+k)%n
			if s[ai] != s[bj] {
				return s[ai] < s[bj]
			}
		}
		return false
	})

	// 构建BWT结果（最后一列）
	result := make([]byte, n)
	for i, idx := range rotations {
		lastCharIndex := (idx + n - 1) % n
		result[i] = s[lastCharIndex]
	}

	return string(result)
}

// BWT逆变换
func inverseBWT(bwt string) string {
	n := len(bwt)
	if n == 0 {
		return ""
	}

	// 初始化索引和计数表
	table := make([][]byte, n)
	for i := range table {
		table[i] = make([]byte, n)
	}

	// 逐步重建矩阵
	for col := n - 1; col >= 0; col-- {
		// 添加BWT列作为第一列
		for i := 0; i < n; i++ {
			table[i][col] = bwt[i]
		}

		// 按行排序
		sort.Slice(table, func(i, j int) bool {
			return string(table[i]) < string(table[j])
		})
	}

	// 查找并返回原始字符串
	for _, row := range table {
		if row[n-1] == '$' {
			return string(row[:n-1])
		}
	}
	return ""
}

func main() {
	input := "bananananananananananananananananana"
	fmt.Printf("原始字符串: %q\n", input)

	// 正向BWT
	bwtResult := BWT(input)
	fmt.Printf("BWT 结果: %q\n", bwtResult) // "annb$aa"

	// 逆向BWT
	recovered := inverseBWT(bwtResult)
	fmt.Printf("恢复的字符串: %q\n", recovered) // "banana"
}
