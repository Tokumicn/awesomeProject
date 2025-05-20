package ai_model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

type GetEmbeddingReq struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type GetEmbeddingResp struct {
	Model           string      `json:"model"`
	Embeddings      [][]float32 `json:"embeddings"`
	TotalDuration   int         `json:"total_duration"`
	LoadDuration    int         `json:"load_duration"`
	PromptEvalCount int         `json:"prompt_eval_count"`
}

func GetEmbedding(ctx context.Context, query string) (*GetEmbeddingResp, error) {
	url := "http://localhost:11434/api/embed" // TODO 抽取配置文件
	req := GetEmbeddingReq{
		Model: "qwen3:0.6b",
		Input: query,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("post embedding http code not 200")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &GetEmbeddingResp{}
	err = json.Unmarshal(respBody, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
