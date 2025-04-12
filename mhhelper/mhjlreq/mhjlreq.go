// 询问梦幻精灵

package mhjlreq

import (
	"context"
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/sagikazarmark/slog-shim"
	"log"
)

const mhjlUrl = "https://xyq.gm.163.com/cgi-bin/csa/csa_sprite.py"

type AnswerResp struct {
	Answer string `json:"answer"`
}

// GetMHJLAnswer 获取梦幻精灵的返回Html div
func GetMHJLAnswer(ctx context.Context, question string) (string, error) {
	resp, err := grequests.Get(mhjlUrl, &grequests.RequestOptions{
		Params: map[string]string{
			"act":          "ask",
			"question":     question,
			"product_name": "xyq",
		}})
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	var anResp AnswerResp
	err = json.Unmarshal(resp.Bytes(), &anResp)
	if err != nil {
		slog.ErrorContext(ctx, "GetMHJLAnswer 获取结果错误 [question: %s] [err: %s]", question, err.Error())
		return "", err
	}

	slog.DebugContext(ctx, "GetMHJLAnswer 获取结果成功 Answer: ", anResp.Answer)
	return anResp.Answer, nil
}
