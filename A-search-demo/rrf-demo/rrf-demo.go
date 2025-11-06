package main

import (
	"fmt"
	"sort"
)

// DocumentScore 表示文档及其得分
type DocumentScore struct {
	DocID string
	Score float64
}

// RRFResult 表示RRF排序结果
type RRFResult struct {
	DocID    string
	RRFScore float64
	Ranks    []int // 在每个列表中的排名
}

// RRFRanker RRF排序器
type RRFRanker struct {
	K float64 // 调和常数，通常取60
}

// NewRRFRanker 创建RRF排序器
func NewRRFRanker(k float64) *RRFRanker {
	return &RRFRanker{K: k}
}

// Fuse 合并多个排名列表
func (r *RRFRanker) Fuse(rankLists [][]string) []RRFResult {
	if len(rankLists) == 0 {
		return nil
	}

	// 收集所有文档ID
	allDocs := make(map[string]bool)
	for _, list := range rankLists {
		for _, docID := range list {
			allDocs[docID] = true
		}
	}

	// 为每个列表构建文档到排名的映射
	listRankMaps := make([]map[string]int, len(rankLists))
	for i, list := range rankLists {
		rankMap := make(map[string]int)
		for rank, docID := range list {
			rankMap[docID] = rank + 1 // 排名从1开始
		}
		listRankMaps[i] = rankMap
	}

	// 计算每个文档的RRF得分
	results := make([]RRFResult, 0, len(allDocs))
	for docID := range allDocs {
		var rrfScore float64
		ranks := make([]int, len(rankLists))

		for i, rankMap := range listRankMaps {
			if rank, exists := rankMap[docID]; exists {
				rrfScore += 1.0 / (r.K + float64(rank))
				ranks[i] = rank
			} else {
				ranks[i] = -1 // 表示文档不在该列表中
			}
		}

		results = append(results, RRFResult{
			DocID:    docID,
			RRFScore: rrfScore,
			Ranks:    ranks,
		})
	}

	// 按RRF得分降序排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].RRFScore > results[j].RRFScore
	})

	return results
}

// FuseWithScores 合并带得分的排名列表
func (r *RRFRanker) FuseWithScores(rankLists [][]DocumentScore) []RRFResult {
	if len(rankLists) == 0 {
		return nil
	}

	// 转换为文档ID列表
	docIDLists := make([][]string, len(rankLists))
	for i, scoredList := range rankLists {
		docIDs := make([]string, len(scoredList))
		for j, ds := range scoredList {
			docIDs[j] = ds.DocID
		}
		docIDLists[i] = docIDs
	}

	return r.Fuse(docIDLists)
}

// AdvancedRRFRanker 支持权重的增强版RRF
type AdvancedRRFRanker struct {
	K       float64
	Weights []float64 // 每个列表的权重
}

// NewAdvancedRRFRanker 创建带权重的RRF排序器
func NewAdvancedRRFRanker(k float64, weights []float64) *AdvancedRRFRanker {
	return &AdvancedRRFRanker{
		K:       k,
		Weights: weights,
	}
}

// Fuse 带权重的RRF合并
func (r *AdvancedRRFRanker) Fuse(rankLists [][]string) []RRFResult {
	if len(rankLists) == 0 {
		return nil
	}

	// 如果没有设置权重，使用默认权重1.0
	weights := r.Weights
	if len(weights) != len(rankLists) {
		weights = make([]float64, len(rankLists))
		for i := range weights {
			weights[i] = 1.0
		}
	}

	allDocs := make(map[string]bool)
	for _, list := range rankLists {
		for _, docID := range list {
			allDocs[docID] = true
		}
	}

	listRankMaps := make([]map[string]int, len(rankLists))
	for i, list := range rankLists {
		rankMap := make(map[string]int)
		for rank, docID := range list {
			rankMap[docID] = rank + 1
		}
		listRankMaps[i] = rankMap
	}

	results := make([]RRFResult, 0, len(allDocs))
	for docID := range allDocs {
		var rrfScore float64
		ranks := make([]int, len(rankLists))

		for i, rankMap := range listRankMaps {
			if rank, exists := rankMap[docID]; exists {
				rrfScore += weights[i] / (r.K + float64(rank))
				ranks[i] = rank
			} else {
				ranks[i] = -1
			}
		}

		results = append(results, RRFResult{
			DocID:    docID,
			RRFScore: rrfScore,
			Ranks:    ranks,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].RRFScore > results[j].RRFScore
	})

	return results
}

func main() {
	// 示例1: 基本RRF排序
	fmt.Println("=== 基本RRF排序示例 ===")

	// 模拟三个不同的排名列表
	list1 := []string{"doc1", "doc2", "doc3", "doc4"}
	list2 := []string{"doc3", "doc1", "doc4", "doc5"}
	list3 := []string{"doc2", "doc5", "doc1", "doc3"}

	rankLists := [][]string{list1, list2, list3}

	// 创建RRF排序器（k=60是常用值）
	ranker := NewRRFRanker(60.0)
	results := ranker.Fuse(rankLists)

	fmt.Println("最终排序结果:")
	for i, result := range results {
		fmt.Printf("%d. 文档%s: RRF得分=%.6f, 排名=%v\n",
			i+1, result.DocID, result.RRFScore, result.Ranks)
	}

	// 示例2: 带权重的RRF排序
	fmt.Println("\n=== 带权重的RRF排序示例 ===")

	weights := []float64{1.0, 2.0, 0.5} // 第二个列表权重最高
	advancedRanker := NewAdvancedRRFRanker(60.0, weights)
	advancedResults := advancedRanker.Fuse(rankLists)

	fmt.Println("带权重的排序结果:")
	for i, result := range advancedResults {
		fmt.Printf("%d. 文档%s: RRF得分=%.6f\n",
			i+1, result.DocID, result.RRFScore)
	}

	// 示例3: 使用带得分的输入
	fmt.Println("\n=== 带得分的RRF排序示例 ===")

	scoredList1 := []DocumentScore{
		{"doc1", 0.9}, {"doc2", 0.8}, {"doc3", 0.7}, {"doc4", 0.6},
	}
	scoredList2 := []DocumentScore{
		{"doc3", 0.95}, {"doc1", 0.85}, {"doc4", 0.75}, {"doc5", 0.65},
	}

	scoredLists := [][]DocumentScore{scoredList1, scoredList2}
	scoredResults := ranker.FuseWithScores(scoredLists)

	fmt.Println("带得分输入的排序结果:")
	for i, result := range scoredResults {
		fmt.Printf("%d. 文档%s: RRF得分=%.6f\n",
			i+1, result.DocID, result.RRFScore)
	}
}
