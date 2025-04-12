package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/textsplitter"
	"log"
	"log/slog"
	"os"
)

var slg slog.Logger

func init() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)
}

func main() {

	file, _ := os.Open("/Users/zhangrui/GolandProjects/awesomeProject/AI-demo/langchain-go-demo/data/shitouji_1_8.txt")
	loader := documentloaders.NewText(file)

	splitter := textsplitter.NewRecursiveCharacter()
	splitter.ChunkSize = 1000  // 单块最大字符数
	splitter.ChunkOverlap = 50 // 块间重叠字符
	docs, _ := loader.LoadAndSplit(context.Background(), splitter)

	// fmt.Println(docs)

	opts := []openai.Option{
		//openai.WithModel("gpt-3.5-turbo-0125"),
		openai.WithEmbeddingModel("text-embedding-v1"),
		openai.WithToken("sk-3fdd61215aa04f29864c3ad7b3a5276b"),
	}
	llm, err := openai.New(opts...)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	embedings, err := llm.CreateEmbedding(ctx, []string{docs[0].PageContent})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(embedings)

	//embeddings := openai.NewEmbeddings(openai.WithToken("YOUR_API_KEY"))

	//drantClient, _ := qdrant.NewClient(qdrant.WithURL("http://localhost:6333"))
	//store := qdrant.NewStore(qdrantClient, "collection-name", embeddings)
	//
	//_, err := store.AddDocuments(context.Background(), docs)
	//if err != nil {
	//	log.Fatal("写入失败:", err)
	//}
}
