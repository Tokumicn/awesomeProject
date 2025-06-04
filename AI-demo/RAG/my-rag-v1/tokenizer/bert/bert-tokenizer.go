package bert

import (
	tokenizer "github.com/Hank-Kuo/go-bert-tokenizer"
)

// BERTTokenizer BERT 分词器
func BERTTokenizer(sentence string) *tokenizer.Encode {

	voc, err := tokenizer.FromFile("./tmp/vocab.txt") // load vocab for vocab file
	if err != nil {
		panic(err)
	}
	tkz := tokenizer.NewFullTokenizer(voc, 128, true)
	encoding := tkz.Tokenize(sentence)

	return encoding
}
