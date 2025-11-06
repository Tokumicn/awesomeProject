package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// 哈夫曼树节点
type Node struct {
	char      rune  // 字符（仅叶子节点有效）
	frequency int   // 频率
	left      *Node // 左子节点
	right     *Node // 右子节点
}

// 优先队列（最小堆）
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// 统计字符频率
func buildFrequencyTable(data string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range data {
		freq[char]++
	}
	return freq
}

// 构建哈夫曼树
func buildHuffmanTree(freqTable map[rune]int) *Node {
	// 创建优先队列
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 添加所有字符作为叶子节点
	for char, frequency := range freqTable {
		heap.Push(&pq, &Node{
			char:      char,
			frequency: frequency,
		})
	}

	// 特殊情况：只有一个字符
	if pq.Len() == 1 {
		left := heap.Pop(&pq).(*Node)
		return &Node{
			frequency: left.frequency,
			left:      left,
		}
	}

	// 构建哈夫曼树
	for pq.Len() > 1 {
		// 取出两个最小频率的节点
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		// 创建新节点（内部节点，char为0）
		parent := &Node{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}

		heap.Push(&pq, parent)
	}

	return heap.Pop(&pq).(*Node)
}

// 生成编码表（递归）
func buildCodeTable(node *Node, code string, table map[rune]string) {
	if node == nil {
		return
	}

	// 叶子节点，存储编码
	if node.left == nil && node.right == nil {
		table[node.char] = code
		return
	}

	// 递归遍历左右子树
	buildCodeTable(node.left, code+"0", table)
	buildCodeTable(node.right, code+"1", table)
}

// 哈夫曼编码
func huffmanEncode(data string) (string, map[rune]string, *Node) {
	if data == "" {
		return "", nil, nil
	}

	// 构建频率表
	freqTable := buildFrequencyTable(data)

	// 构建哈夫曼树
	tree := buildHuffmanTree(freqTable)

	// 生成编码表
	codeTable := make(map[rune]string)
	buildCodeTable(tree, "", codeTable)

	// 编码数据
	var encoded strings.Builder
	for _, char := range data {
		encoded.WriteString(codeTable[char])
	}

	return encoded.String(), codeTable, tree
}

// 哈夫曼解码
func huffmanDecode(encoded string, tree *Node) string {
	if encoded == "" || tree == nil {
		return ""
	}

	var decoded strings.Builder
	current := tree

	for _, bit := range encoded {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		// 到达叶子节点
		if current.left == nil && current.right == nil {
			decoded.WriteRune(current.char)
			current = tree // 回到根节点
		}
	}

	return decoded.String()
}

// 打印编码表（按频率排序）
func printCodeTable(codeTable map[rune]string, freqTable map[rune]int) {
	fmt.Println("哈夫曼编码表:")
	fmt.Println("字符\t频率\t编码")

	// 创建排序列表
	type CharInfo struct {
		char rune
		freq int
		code string
	}

	var list []CharInfo
	for char, code := range codeTable {
		list = append(list, CharInfo{
			char: char,
			freq: freqTable[char],
			code: code,
		})
	}

	// 按频率降序排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].freq > list[j].freq
	})

	for _, item := range list {
		charStr := string(item.char)
		if item.char == ' ' {
			charStr = "[空格]"
		} else if item.char == '\n' {
			charStr = "[换行]"
		}
		fmt.Printf("%s\t%d\t%s\n", charStr, item.freq, item.code)
	}
}

// 计算压缩率
func calculateCompressionRatio(original, encoded string, codeTable map[rune]string) {
	originalBits := len(original) * 8 // 假设原始是ASCII，每个字符8位
	encodedBits := len(encoded)

	// 理论平均编码长度
	totalFreq := 0
	averageLength := 0.0
	freqTable := buildFrequencyTable(original)
	for char, freq := range freqTable {
		totalFreq += freq
		averageLength += float64(freq) * float64(len(codeTable[char]))
	}
	averageLength /= float64(totalFreq)

	fmt.Printf("\n压缩统计:\n")
	fmt.Printf("原始数据大小: %d 字符 × 8 位 = %d 位\n", len(original), originalBits)
	fmt.Printf("编码后大小: %d 位\n", encodedBits)
	fmt.Printf("压缩率: %.2f%%\n", float64(encodedBits)/float64(originalBits)*100)
	fmt.Printf("平均编码长度: %.2f 位/字符\n", averageLength)
	fmt.Printf("空间节省: %.2f%%\n", (1-float64(encodedBits)/float64(originalBits))*100)
}

func main() {
	testCases := []string{
		"this is an example for huffman encoding",
		"abcdefghijklmnopqrstuvwxyz",
		"aaaaabbbbbccccccdddddddeeeeeee",
		"hello world",
		"哈夫曼编码测试", // 测试中文字符
		"a1n2b1$1a2",
	}

	for i, data := range testCases {
		fmt.Printf("=== 测试用例 %d ===\n", i+1)
		fmt.Printf("原始数据: %s\n", data)

		// 编码
		encoded, codeTable, tree := huffmanEncode(data)
		freqTable := buildFrequencyTable(data)
		printCodeTable(codeTable, freqTable)

		fmt.Printf("\n编码结果: %s\n", encoded)

		// 解码
		decoded := huffmanDecode(encoded, tree)
		fmt.Printf("解码结果: %s\n", decoded)

		// 压缩统计
		calculateCompressionRatio(data, encoded, codeTable)

		fmt.Printf("\n%s\n\n", strings.Repeat("-", 50))
	}

	// 详细示例分析
	detailedExample()
}

func detailedExample() {
	fmt.Println("=== 详细示例分析 ===")
	data := "a1n2b1$1a2"
	fmt.Printf("原始字符串: %s\n\n", data)

	// 统计频率
	freqTable := buildFrequencyTable(data)
	fmt.Println("字符频率统计:")
	for char, freq := range freqTable {
		fmt.Printf("'%c': %d\n", char, freq)
	}

	// 构建哈夫曼树和编码
	encoded, codeTable, tree := huffmanEncode(data)

	fmt.Printf("\n哈夫曼编码表:\n")
	for char, code := range codeTable {
		fmt.Printf("'%c' → %s\n", char, code)
	}

	fmt.Printf("\n编码过程:\n")
	for _, char := range data {
		fmt.Printf("'%c'(%s) ", char, codeTable[char])
	}
	fmt.Printf("\n编码结果: %s\n", encoded)

	// 验证编码解码
	decoded := huffmanDecode(encoded, tree)
	fmt.Printf("解码验证: %s\n", decoded)

	fmt.Printf("编解码是否一致: %t\n", data == decoded)
}
