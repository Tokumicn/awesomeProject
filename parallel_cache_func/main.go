package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient represents a Redis client with a name for identification.
type RedisClient struct {
	client *redis.Client
	name   string
}

// NewRedisClient creates a new Redis client.
func NewRedisClient(address, password string, db int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{client: client}, nil
}

// GetFromCache queries the cache and returns the value.
func (rc *RedisClient) GetFromCache(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist in %s", key, rc.name)
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func main() {
	// Initialize Redis clients
	client1, err := NewRedisClient("localhost:6379", "", 0)
	if err != nil {
		fmt.Println("Failed to create Redis client 1:", err)
		return
	}
	defer client1.client.Close()

	client2, err := NewRedisClient("localhost:6380", "", 0)
	if err != nil {
		fmt.Println("Failed to create Redis client 2:", err)
		return
	}
	defer client2.client.Close()

	clients := []*RedisClient{{client: client1.client, name: "Redis1"}, {client: client2.client, name: "Redis2"}}

	// Set some test data
	client1.client.Set(context.Background(), "testKey", "valueFromRedis1", 0).Err()
	client2.client.Set(context.Background(), "testKey", "valueFromRedis2", 0).Err()

	// Query the cache
	value, err := QueryLevelCache(context.Background(), clients, "testKey")
	if err != nil {
		fmt.Println("Error querying cache:", err)
		return
	}
	fmt.Println("Value from cache:", value)
}

// QueryLevelCache
// golang实现一个函数，查询两级Redis缓存，要求并行查询谁先返回用谁，两处缓存结构一致
func QueryLevelCache(ctx context.Context, clients []*RedisClient, key string) (string, error) {
	type result struct {
		value string
		err   error
	}

	resultChan := make(chan result, len(clients))
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	for _, client := range clients {
		go func(client *RedisClient) {
			val, err := client.GetFromCache(ctx, key)
			select {
			case <-ctx.Done():
				return
			default:
				resultChan <- result{value: val, err: err}
			}
		}(client)
	}

	for i := 0; i < cap(resultChan); i++ {
		select {
		case res := <-resultChan:
			if res.err == nil {
				return res.value, nil
			}
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	return "", fmt.Errorf("all cache queries failed")
}

// QueryLevelCacheV2
// golang实现一个函数，查询两级Redis缓存，要求并行查询谁先返回用谁，两处缓存结构一致。
// 但是如果第一个返回是除了redis.Nil之外的错误，则等待第二个返回结果再结束
func QueryLevelCacheV2(ctx context.Context, clients []*RedisClient, key string) (string, error) {
	type result struct {
		value string
		err   error
	}

	resultChan := make(chan result, len(clients))
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	for _, client := range clients {
		go func(client *RedisClient) {
			val, err := client.GetFromCache(ctx, key)
			select {
			case <-ctx.Done():
				return
			default:
				resultChan <- result{value: val, err: err}
			}
		}(client)
	}

	firstError := false
	for i := 0; i < cap(resultChan); i++ {
		select {
		case res := <-resultChan:
			if res.err == nil {
				return res.value, nil
			}
			if !firstError && res.err != redis.Nil {
				firstError = true
				continue // Wait for second response
			}
			return "", res.err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	return "", fmt.Errorf("all cache queries failed")
}
