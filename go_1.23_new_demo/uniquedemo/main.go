package main

import (
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"unique"
)

var words []string

func main() {
	const (
		nWords    = 10000
		nDistinct = 100
		wordLen   = 40
	)
	generate := wordGen(nDistinct, wordLen)
	memBefore := getAlloc()

	words = make([]string, nWords)
	for i := range words {
		words[i] = generate()
	}

	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	fmt.Printf("Slice Memory used: %dKB\n", memUsed/1024)

	memBeforeV2 := getAlloc()

	// Go 1.23 版本新特性
	// 字符串复用节省空间
	wordsV2 := make([]unique.Handle[string], nWords)
	for i := range wordsV2 {
		wordsV2[i] = unique.Make(generate())
	}

	memAfterV2 := getAlloc()
	memUsedV2 := memAfterV2 - memBeforeV2
	fmt.Printf("Unique Memory used: %dKB\n", memUsedV2/1024)

}

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func wordGen(nDistinct, wordLen int) func() string {
	vocab := make([]string, nDistinct)
	for i := range vocab {
		word := randomString(wordLen)
		vocab[i] = word
	}

	return func() string {
		word := vocab[rand.Intn(len(vocab))]
		return strings.Clone(word)
	}
}

func randomString(n int) string {
	const letters = "eddycjyabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; {
		b := make([]byte, 1)
		if _, err := crand.Read(b); err != nil {
			panic(err)
		}
		ret[i] = letters[int(b[0])%len(letters)]
		i++
	}
	return string(ret)
}
