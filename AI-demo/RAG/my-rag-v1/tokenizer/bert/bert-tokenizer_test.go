package bert

import (
	"fmt"
	"testing"
)

func TestBertTokenizerDemo(t *testing.T) {

	sentence := "我爱北京天安门"
	encode := BERTTokenizer(sentence)

	fmt.Println(encode.Text)
	fmt.Println(encode.Tokens)
	fmt.Println(encode.TokenIDs)
	fmt.Println(encode.MaskIDs)
	fmt.Println(encode.TypeIDs)
}
