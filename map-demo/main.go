package main

import (
	"awesomeProject/map-demo/shardmap"
	"fmt"
	"sync"
)

func main() {
	mapForNil()
	//demoShardMap()
	//demoSyncMap()
}

func mapForNil() {
	m := map[string]interface{}{}
	m = nil

	val, ok := m["name"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Printf("map nit exist key name\n")
	}

	val2 := m["name"]
	fmt.Println(val2)
}

// shardmap 使用示例
func demoShardMap() {
	sm := shardmap.NewShardedMap()
	var wg sync.WaitGroup

	// 并发写入
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sm.Set(fmt.Sprintf("key%d", i), i)
		}
	}()

	// 并发读取
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if val, ok := sm.Get(fmt.Sprintf("key%d", i)); ok {
				fmt.Printf("Got: %v\n", val)
			}
		}
	}()

	wg.Wait()
}

// sync map 使用示例
func demoSyncMap() {
	var m sync.Map
	var wg sync.WaitGroup

	// 写入数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			m.Store(i, i*i)
		}
	}()

	// 读取数据
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if val, ok := m.Load(i); ok {
				fmt.Printf("Key: %d, Value: %v\n", i, val)
			}
		}
	}()

	// 使用Range遍历
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Range(func(k, v interface{}) bool {
			fmt.Printf("Range: %v -> %v\n", k, v)
			return true // 继续遍历
		})
	}()

	wg.Wait()
}
