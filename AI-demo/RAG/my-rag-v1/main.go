package main

import (
	"ai-demo/RAG/my-rag-v1/ai_model"
	"ai-demo/RAG/my-rag-v1/tokenizer"
	"ai-demo/RAG/my-rag-v1/tokenizer/jieba_spliter"
	"ai-demo/RAG/my-rag-v1/vector_db"
	"context"
	"fmt"
	"github.com/pgvector/pgvector-go"
	"github.com/tmc/langchaingo/textsplitter"
	"log/slog"
	"os"
)

func init() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)

	_, err := vector_db.NewDB(context.TODO())
	if err != nil {
		panic(err)
	}
}

func main() {

	ctx := context.TODO()
	defer func() {
		if err := recover(); err != nil {
			slog.ErrorContext(ctx, "panic: ", err)
		}
	}()

	// 明天的TODO: 1-优化分词，看向量搜索的效果  2-将向量知识拼接prompt，让模型回答问题

	// TODO 上传文件

	// 构建知识库
	//err := buildKnowledge(ctx)
	//if err != nil {
	//	slog.ErrorContext(ctx, "buildKnowledge err: ", err)
	//	return
	//}

	// 查询知识库
	queries := []string{
		//"鱼钩上不放鱼饵能钓到吗？",
		//"鱼能生吃吗?",
		//"水能燃烧?",
		//"六物是什么?",
		//"什么是分？什么是命？",
		"八卦是什么？",
		//"八卦",
	}

	for _, q := range queries {
		knowledge, err := searchKnowledge(ctx, q)
		if err != nil {
			slog.ErrorContext(ctx, "searchKnowledge err: ", err)
			continue
		}

		fmt.Println("Q: ", q)
		for _, k := range knowledge {
			fmt.Println("A: ", k.Content)
		}
		fmt.Println("------------------------------")
	}
}

func searchKnowledge(ctx context.Context, query string) ([]vector_db.DocVector, error) {
	// 对问题分词

	splitter := tokenizer.NewSplitter(tokenizer.Options{
		ChunkSize:    256,
		ChunkOverlap: 25,
		CutType:      jieba_spliter.CutModeAccurate,
		SplitterType: tokenizer.LangChain,
	})
	strings := splitter.Split(query)

	allDocs := make([]vector_db.DocVector, 0)

	for _, subQ := range strings {
		em, err := ai_model.GetEmbedding(ctx, subQ)
		if err != nil {
			return nil, err
		}

		docVec := vector_db.DocVector{}
		similarQuery, err := docVec.SimilarQuery(ctx, pgvector.NewVector(em.Embeddings[0]))
		if err != nil {
			return nil, err
		}

		allDocs = append(allDocs, similarQuery...)
	}

	return allDocs, nil
}

func buildKnowledge(ctx context.Context) error {
	// 读取文件
	contents, err := readFile(ctx, "TODO")
	if err != nil {
		slog.ErrorContext(ctx, "readFile err: ", err)
		return err
	}

	// 生成embedding
	err = buildEmbedding(ctx, contents)
	if err != nil {
		slog.ErrorContext(ctx, "buildEmbedding err: ", err)
		return err
	}
	return nil
}

func readFile(ctx context.Context, filePath string) ([]string, error) {
	// 读取文件 /data/渔樵对.txt
	//exePath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//binDir := filepath.Dir(exePath)
	//filePath := fmt.Sprintf("%s/data/渔樵对.txt", binDir)
	bytes, err := os.ReadFile("./data/渔樵对.txt")
	if err != nil {
		slog.ErrorContext(ctx, "os.ReadFile err: ", err)
		return nil, err
	}

	// 初始化分词器
	split := textsplitter.NewRecursiveCharacter()
	split.Separators = []string{"问："}
	split.ChunkSize = 512   // 最大长度
	split.ChunkOverlap = 50 // 块大小
	contents, err := split.SplitText(string(bytes))
	if err != nil {
		slog.ErrorContext(ctx, "split.SplitText err: ", err)
		return nil, err
	}
	return contents, nil
}

func buildEmbedding(ctx context.Context, contents []string) error {
	var (
		docID   uint = 1
		blockID uint = 1
	)
	// 切分文件  普通切分  语义切分  大模型切分
	for _, content := range contents {
		emResp, err := ai_model.GetEmbedding(ctx, content)
		if err != nil {
			slog.ErrorContext(ctx, "ai_model.GetEmbedding err: ", err)
			return err
		}

		if emResp == nil || len(emResp.Embeddings) <= 0 {
			slog.ErrorContext(ctx, "embeddingResp is nil content: ", content)
			return err
		}

		temp := emResp.Embeddings[0]

		tempDoc := vector_db.DocVector{
			DocID:       docID,
			BlockID:     blockID,
			Content:     content,
			Embedding:   pgvector.NewVector(temp),
			TrainStatus: 0,
		}

		err = tempDoc.Create()
		if err != nil {
			slog.ErrorContext(ctx, "tempDoc.Create err: ", err)
			continue
		}
	}

	return nil
}
