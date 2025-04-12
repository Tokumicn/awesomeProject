package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	g := new(singleflight.Group)
	dataKey := "getData"

	go func() {
		v1, _, shared := g.Do(dataKey, func() (interface{}, error) {
			ret := getData(1)
			return ret, nil
		})
		fmt.Printf("1st call: v1:%v, shared:%v\n", v1, shared)
	}()

	time.Sleep(1 * time.Second)

	v2, _, shared := g.Do(dataKey, func() (interface{}, error) {
		ret := getData(2)
		return ret, nil
	})
	fmt.Printf("2nd call: v2:%v, shared:%v\n", v2, shared)
}

func getData(id int64) string {
	fmt.Println("query...")
	time.Sleep(10 * time.Second)
	return "hello tom."
}

//func getDataSingleFlight(key string) (interface{}, error) {
//	g := new(singleflight.Group)
//	v, err, _ := g.Do(key, func() (interface{}, error) {
//		// 查缓存
//		data, err := getDataFromCache(key)
//		if err == nil {
//			return data, nil
//		}
//		if err == errNotFound {
//			// 查DB
//			data, err := getDataFromDB(key)
//			if err == nil {
//				setCache(data) // 设置缓存
//				return data, nil
//			}
//			return nil, err
//		}
//		return nil, err // 缓存出错直接返回，防止灾难传递至DB
//	})
//
//	if err != nil {
//		return nil, err
//	}
//	return v, nil
//}
