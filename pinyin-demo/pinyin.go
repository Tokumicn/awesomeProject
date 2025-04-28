package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "汉字转拼音"
	py := pinyin.Pinyin(hans, pinyin.NewArgs())
	fmt.Println(py) // 输出: [[han] [zi] [zhuan] [pin] [yin]]

	// 带音调
	args := pinyin.NewArgs()
	args.Style = pinyin.Tone // 设置拼音风格为带音调
	py = pinyin.Pinyin("汉字", args)
	fmt.Println(py) // 输出: [[hàn] [zì]]

	// 数字表示音调
	args.Style = pinyin.Tone3 // 数字音调在韵母后  Tone2: 数字音调在声母后s
	py = pinyin.Pinyin("汉字", args)
	fmt.Println(py) // 输出: [[han4] [zi4]]

	// 声母（首字母）
	args.Style = pinyin.Initials
	py = pinyin.Pinyin("汉字", args)
	fmt.Println(py) // 输出: [[h] [z]]

	// 仅韵母
	args.Style = pinyin.Finals
	py = pinyin.Pinyin("汉字", args)
	fmt.Println(py) // 输出: [[an] [i]]

	// 带 ü 的拼音
	args.Style = pinyin.Tone3
	py = pinyin.Pinyin("绿", args)
	fmt.Println(py) // 输出: [[lü4]]

	// 获取拼音首字母如果只想提取拼音的首字母，可以使用 pinyin.FirstLetter：
	pys := pinyin.LazyPinyin("汉字", pinyin.NewArgs())
	fmt.Println(pys) // 输出: [h z]

}
