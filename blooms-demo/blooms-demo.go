package main

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// 创建一个布隆过滤器，使用 1000 个 bit 位，3 个哈希函数
	filter := bloom.New(1000, 3)
	// 添加数据
	filter.Add([]byte("hello"))
	filter.Add([]byte("world"))
	// 查询数据
	fmt.Println(filter.Test([]byte("hello")))  // true
	fmt.Println(filter.Test([]byte("world")))  // true
	fmt.Println(filter.Test([]byte("golang"))) // false（一定正确）

	// 计算误判率
	n := uint(1000)                        // 预期存储的元素个数
	p := 0.01                              // 误判率
	m, k := bloom.EstimateParameters(n, p) // 计算最佳位数组大小
	fmt.Println(m, k)                      // 输出建议的位数组大小  m:比特位个数    k:哈希函数个数
	// ActualfpRate := bloom.EstimateFalsePositiveRate(m, k, n)
}
