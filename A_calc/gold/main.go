package main

import (
	"fmt"
	"net/http"
)

/*
Q1：黄金美元/盎司如何换算成人民币/克?
美元/盎司对应多少元/克？ 需要考虑两个因素：汇率和单位换算。下面是换算步骤：

汇率换算：首先，需要知道当前的美元对人民币的汇率。
单位换算：其次，需要知道1盎司等于多少克。1盎司大约等于31.1035克。
计算：然后，将美元/盎司的价格乘以汇率，再除以盎司到克的换算系数，得到人民币/克的价格。
具体公式如下：

人民币/克的金价=(美元盎司的金价 × 美元对人民币汇率 / 31.1035)

比如：3000美元每盎司是多少钱一克？
计算：假设当前的汇率是1美元兑换7人民币，那么换算成人民币/克的价格为：人民币/克=(3000×7/31.1035)≈675.16

Q2：一盎司黄金是多少克?
1盎司黄金等于31.1034768克，一般大约等于31.1035克

不同国家和地区在黄金计量标准上存在一些差异，以下是一些主要的区别：

金衡盎司（Troy Ounce）
在国际市场上，黄金计量主要使用金衡盎司作为单位，每金衡盎司等于31.1035克。这是全球黄金交易中最常用的计量单位，尤其在国际现货市场和期货市场中广泛使用。

克（Gram）
在中国，黄金计量主要使用克和千克。中国的黄金市场自2002年开放以来，逐渐形成了以克为单位的交易习惯。

托拉（Tola）
在印度，黄金的传统计量单位是托拉，每托拉等于11.6638克。此外，印度也使用克和盎司作为计量单位。

克希（Kilo）
在中东地区，尤其是阿拉伯国家，黄金计量通常使用克希和盎司。1克希等于1000克。
*/
const (
	// 1盎司黄金等于31.1034768克
	ounceToGram = 31.1034768
	// 1美元等于多少人民币
	dollarToRMB = 7
)

// GetExchangeRate
// 获取汇率 curl --location 'https://www.jins.gold/api/exchange-rate'
func GetExchangeRate() (float64, error) {

	resp, err := http.Get("https://www.jins.gold/api/exchange-rate")
	if err != nil {
		return 0, err
	}
	fmt.Println(resp)

	return 7, nil
}

func main() {

	rate, err := GetExchangeRate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rate)

	// 3000美元每盎司是多少钱一克？
	price := 3000.0
	// 计算：假设当前的汇率是1美元兑换7人民币，那么换算成人民币/克的价格为：人民币/克=(3000×7/31.1035)≈675.16
	priceRMB := (price * rate / ounceToGram)
	fmt.Println(priceRMB)
}
