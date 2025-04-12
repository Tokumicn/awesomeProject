package shardmap

// 分片锁map 读写性能均衡，适合写多场景
// 1. 读写锁保护，分场景应对读写情况
// 2. 分片内维护各自读写锁
// 3. 对key进行hash分片路由

import (
	"hash/fnv"
	"sync"
)

// 定义分片数量（建议为2的幂）
const shardCount = 16

// 分片结构体
type Shard struct {
	sync.RWMutex
	data map[string]interface{}
}

// 分片Map
type ShardedMap []*Shard

// 创建新分片Map
func NewShardedMap() ShardedMap {
	sm := make(ShardedMap, shardCount)
	for i := 0; i < shardCount; i++ {
		sm[i] = &Shard{data: make(map[string]interface{})}
	}
	return sm
}

// 哈希函数（示例：FNV算法）
func (sm ShardedMap) getShardIndex(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % uint32(shardCount)
}

// 获取分片
func (sm ShardedMap) getShard(key string) *Shard {
	index := sm.getShardIndex(key)
	return sm[index]
}

// Set 写操作
func (sm ShardedMap) Set(key string, value interface{}) {
	shard := sm.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.data[key] = value
}

// Get 读操作
func (sm ShardedMap) Get(key string) (interface{}, bool) {
	shard := sm.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.data[key]
	return val, ok
}

// Delete 删除操作
func (sm ShardedMap) Delete(key string) {
	shard := sm.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.data, key)
}

//// 示例使用
//func main() {
//	sm := NewShardedMap()
//	var wg sync.WaitGroup
//
//	// 并发写入
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 1000; i++ {
//			sm.Set(fmt.Sprintf("key%d", i), i)
//		}
//	}()
//
//	// 并发读取
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for i := 0; i < 1000; i++ {
//			if val, ok := sm.Get(fmt.Sprintf("key%d", i)); ok {
//				fmt.Printf("Got: %v\n", val)
//			}
//		}
//	}()
//
//	wg.Wait()
//}
