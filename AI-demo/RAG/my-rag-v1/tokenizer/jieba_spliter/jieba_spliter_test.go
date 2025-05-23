package jieba_spliter

import (
	"fmt"
	"testing"
)

func TestJiebaSplitter(t *testing.T) {

	splitter := NewSplitter(true)
	fmt.Print("【全模式】：")
	splitter.PrintCh(splitter.segmenter.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	splitter.PrintCh(splitter.segmenter.Cut("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	splitter.PrintCh(splitter.segmenter.Cut("我来到北京清华大学", true))

	fmt.Print("【新词识别】：")
	splitter.PrintCh(splitter.segmenter.Cut("他来到了网易杭研大厦", true))

	fmt.Print("【搜索引擎模式】：")
	splitter.PrintCh(splitter.segmenter.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))

	fmt.Println("================================================")

	fmt.Print("【全模式】：")
	splitter.Print(splitter.Split(CutModeAll, "我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	splitter.Print(splitter.Split(CutModeAccurate, "我来到北京清华大学"))

	fmt.Print("【新词识别】：")
	splitter.Print(splitter.Split(CutModeAccurateNew, "我来到北京清华大学"))

	fmt.Print("【新词识别】：")
	splitter.Print(splitter.Split(CutModeAccurateNew, "他来到了网易杭研大厦"))

	fmt.Print("【搜索引擎模式】：")
	splitter.Print(splitter.Split(CutModeSearch, "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"))

}
