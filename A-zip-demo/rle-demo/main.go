package main

import (
	"fmt"
	"strconv"
	"strings"
)

// RLEEncode 游程编码
func RLEEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1
	currentChar := rune(input[0])

	for i := 1; i < len(input); i++ {
		if rune(input[i]) == currentChar {
			count++
		} else {
			// 写入前一个字符的编码
			if count > 1 {
				result.WriteString(fmt.Sprintf("%c%d", currentChar, count))
			} else {
				result.WriteByte(byte(currentChar))
			}
			currentChar = rune(input[i])
			count = 1
		}
	}

	// 处理最后一个字符序列
	if count > 1 {
		result.WriteString(fmt.Sprintf("%c%d", currentChar, count))
	} else {
		result.WriteByte(byte(currentChar))
	}

	return result.String()
}

// RLEDecode 游程解码
func RLEDecode(encoded string) string {
	var result strings.Builder
	i := 0

	for i < len(encoded) {
		currentChar := rune(encoded[i])

		// 检查后面是否有数字（重复次数）
		if i+1 < len(encoded) {
			j := i + 1
			// 提取所有连续的数字
			for j < len(encoded) && encoded[j] >= '0' && encoded[j] <= '9' {
				j++
			}

			if j > i+1 { // 有数字存在
				count, _ := strconv.Atoi(encoded[i+1 : j])
				result.WriteString(strings.Repeat(string(currentChar), count))
				i = j
				continue
			}
		}

		// 没有数字，直接写入字符
		result.WriteByte(byte(currentChar))
		i++
	}

	return result.String()
}

// 改进版本：处理单个字符和特殊字符
func RLEEncodeImproved(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1

	for i := 1; i <= len(input); i++ {
		if i < len(input) && input[i] == input[i-1] {
			count++
		} else {
			// 总是写入计数，便于统一解码
			result.WriteString(fmt.Sprintf("%d%c", count, input[i-1]))
			count = 1
		}
	}

	return result.String()
}

func RLEDecodeImproved(encoded string) string {
	var result strings.Builder
	i := 0

	for i < len(encoded) {
		// 提取数字部分
		j := i
		for j < len(encoded) && encoded[j] >= '0' && encoded[j] <= '9' {
			j++
		}

		if j > i { // 成功提取到数字
			count, _ := strconv.Atoi(encoded[i:j])
			if j < len(encoded) {
				char := encoded[j]
				result.WriteString(strings.Repeat(string(char), count))
				i = j + 1
			} else {
				break
			}
		} else {
			// 格式错误，跳过
			i++
		}
	}

	return result.String()
}

func main() {
	testCases := []string{
		"AAAAABBBCCDAA", // 典型用例
		"ABC",           // 无重复字符
		"A",             // 单个字符
		"AABBCC",        // 简单重复
		"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", // 复杂用例
		"112233", // 数字字符
		"annb$aa",
	}

	fmt.Println("=== 基本版本测试 ===")
	for _, test := range testCases {
		encoded := RLEEncode(test)
		decoded := RLEDecode(encoded)
		fmt.Printf("原始: %-50s → 编码: %-30s → 解码: %s\n",
			test, encoded, decoded)
	}

	fmt.Println("\n=== 改进版本测试 ===")
	for _, test := range testCases {
		encoded := RLEEncodeImproved(test)
		decoded := RLEDecodeImproved(encoded)
		fmt.Printf("原始: %-50s → 编码: %-30s → 解码: %s\n",
			test, encoded, decoded)
	}

	// 压缩率计算示例
	fmt.Println("\n=== 压缩率分析 ===")
	original := "AAAAABBBCCDAA"
	encoded := RLEEncode(original)
	compressionRatio := float64(len(encoded)) / float64(len(original)) * 100
	fmt.Printf("原始数据: %s (%d字节)\n", original, len(original))
	fmt.Printf("编码后: %s (%d字节)\n", encoded, len(encoded))
	fmt.Printf("压缩率: %.1f%%\n", compressionRatio)
}
