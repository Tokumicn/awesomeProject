package main

import (
	"fmt"
	"strings"
)

// TreeNode 表示树节点结构
type TreeNode struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	ParentID int        `json:"parent_id"`
	Sub      []TreeNode `json:"sub,omitempty"`
}

// BuildTree 递归构造树形结构
func buildSubTree(parentID int, group map[int][]TreeNode) []TreeNode {
	nodes := group[parentID]
	for i := range nodes {
		nodes[i].Sub = buildSubTree(nodes[i].ID, group)
	}
	return nodes
}

func BuildTree(nodes []TreeNode) []TreeNode {
	// 先分组：parent_id -> []TreeNode
	group := make(map[int][]TreeNode)
	for _, node := range nodes {
		group[node.ParentID] = append(group[node.ParentID], node)
	}
	return buildSubTree(0, group)
}

// PrintTree 以树形结构打印节点
func PrintTree(nodes []TreeNode, level int) {
	prefix := strings.Repeat("  ", level)
	for _, node := range nodes {
		fmt.Printf("%s├─ %s (ID: %d)\n", prefix, node.Name, node.ID)
		if len(node.Sub) > 0 {
			PrintTree(node.Sub, level+1)
		}
	}
}

func main() {
	// 示例数据
	nodes := []TreeNode{
		{ID: 1, Name: "根节点1", ParentID: 0},
		{ID: 2, Name: "子节点1-1", ParentID: 1},
		{ID: 3, Name: "子节点1-2", ParentID: 1},
		{ID: 4, Name: "子节点1-1-1", ParentID: 2},
		{ID: 5, Name: "根节点2", ParentID: 0},
		{ID: 6, Name: "子节点2-1", ParentID: 5},
		{ID: 7, Name: "子节点2-2", ParentID: 5},
		{ID: 8, Name: "子节点2-2-1", ParentID: 7},
		{ID: 9, Name: "子节点2-2-1-1", ParentID: 8},
	}

	// 构建树形结构
	tree := BuildTree(nodes)

	// 打印树形结构
	fmt.Println("树形结构：")
	PrintTree(tree, 0)
}
