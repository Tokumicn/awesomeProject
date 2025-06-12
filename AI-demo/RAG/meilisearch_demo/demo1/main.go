// In the command line:
// go get -u github.com/meilisearch/meilisearch-go

// In your .go file:
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.New("http://localhost:7700", meilisearch.WithAPIKey("LYBqMzhd2P_CSqmd9m1vgi9wHTREgQkKDeHwn9Hmx-s"))

	// 首次使用即可--构建索引数据
	// err := buildIndex(client)

	search, err := client.Index("movies").Search("botman", &meilisearch.SearchRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(search)
}

func buildIndex(client meilisearch.ServiceManager) error {
	jsonFile, _ := os.Open("movies.json")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var movies []map[string]interface{}
	err := json.Unmarshal(byteValue, &movies)
	if err != nil {
		panic(err)
	}

	_, err = client.Index("movies").AddDocuments(movies)
	if err != nil {
		panic(err)
	}
	return err
}
