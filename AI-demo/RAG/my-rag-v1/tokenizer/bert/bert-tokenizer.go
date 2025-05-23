package bert

import (
	"fmt"
	tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)

// BERTTokenizer BERT 分词器
func BERTTokenizer(sentence string) {
	sentence = "我爱北京天安门"

	voc, err := tokenizer.FromFile("./tmp/vocab.txt") // load vocab for vocab file
	if err != nil {
		panic(err)
	}
	tkz := tokenizer.NewFullTokenizer(voc, 128, true)
	encoding := tkz.Tokenize(sentence)
	fmt.Println(encoding.Text)
	fmt.Println(encoding.Tokens)
	fmt.Println(encoding.TokenIDs)
	fmt.Println(encoding.MaskIDs)
	fmt.Println(encoding.TypeIDs)
}
