package main

import (
	"fmt"
	"strings"
)

// 八卦定义
type Gua struct {
	Name   string
	Symbol string
	Number int
	Yao    [3]int // 爻表示：0-阴，1-阳
}

// 六十四卦定义
type Hexagram struct {
	Name     string
	Symbol   string
	UpperGua *Gua
	LowerGua *Gua
	Yao      [6]int // 从初爻到上爻
}

// 八字信息
type BaZi struct {
	YearZhi  string // 年支
	MonthZhi string // 月支
	DayZhi   string // 日支
	HourZhi  string // 时支
	Gender   string // 性别
}

// 地支到数字的映射
var DiZhiToNumber = map[string]int{
	"子": 1, "丑": 2, "寅": 3, "卯": 4, "辰": 5, "巳": 6,
	"午": 7, "未": 8, "申": 9, "酉": 10, "戌": 11, "亥": 12,
}

// 八卦映射
var BaGuaMap = map[int]*Gua{
	1: {Name: "乾", Symbol: "☰", Number: 1, Yao: [3]int{1, 1, 1}},
	2: {Name: "兑", Symbol: "☱", Number: 2, Yao: [3]int{1, 1, 0}},
	3: {Name: "离", Symbol: "☲", Number: 3, Yao: [3]int{1, 0, 1}},
	4: {Name: "震", Symbol: "☳", Number: 4, Yao: [3]int{0, 0, 1}},
	5: {Name: "巽", Symbol: "☴", Number: 5, Yao: [3]int{0, 1, 1}},
	6: {Name: "坎", Symbol: "☵", Number: 6, Yao: [3]int{0, 1, 0}},
	7: {Name: "艮", Symbol: "☶", Number: 7, Yao: [3]int{1, 0, 0}},
	8: {Name: "坤", Symbol: "☷", Number: 8, Yao: [3]int{0, 0, 0}},
}

// 六十四卦映射（简略版，实际应有64个）
var Gua64Map = map[string]*Hexagram{
	// 乾宫八卦
	"乾为天":  {Name: "乾为天", Symbol: "䷀", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 1, 1, 1}},
	"天风姤":  {Name: "天风姤", Symbol: "䷫", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 1, 1, 1}},
	"天山遁":  {Name: "天山遁", Symbol: "䷠", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 1, 1, 1}},
	"天地否":  {Name: "天地否", Symbol: "䷋", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 1, 1, 1}},
	"风地观":  {Name: "风地观", Symbol: "䷓", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 1, 1}},
	"山地剥":  {Name: "山地剥", Symbol: "䷖", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 0, 1}},
	"火地晋":  {Name: "火地晋", Symbol: "䷢", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 1, 0}},
	"火天大有": {Name: "火天大有", Symbol: "䷍", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 0, 1, 0}},

	// 兑宫八卦
	"兑为泽":  {Name: "兑为泽", Symbol: "䷹", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 1, 1, 0}},
	"泽水困":  {Name: "泽水困", Symbol: "䷮", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 1, 1, 0}},
	"泽地萃":  {Name: "泽地萃", Symbol: "䷬", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 1, 1, 0}},
	"泽山咸":  {Name: "泽山咸", Symbol: "䷞", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 1, 1, 0}},
	"水山蹇":  {Name: "水山蹇", Symbol: "䷦", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 0, 1, 0}},
	"地山谦":  {Name: "地山谦", Symbol: "䷎", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 0, 0, 0}},
	"雷山小过": {Name: "雷山小过", Symbol: "䷽", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 0, 0, 1}},
	"雷泽归妹": {Name: "雷泽归妹", Symbol: "䷵", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 0, 0, 1}},

	// 离宫八卦
	"离为火":  {Name: "离为火", Symbol: "䷝", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 1, 0, 1}},
	"火山旅":  {Name: "火山旅", Symbol: "䷷", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[7], Yao: [6]int{0, 0, 1, 1, 0, 1}},
	"火风鼎":  {Name: "火风鼎", Symbol: "䷱", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 1, 0, 1}},
	"火水未济": {Name: "火水未济", Symbol: "䷿", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 1, 0, 1}},
	"山水蒙":  {Name: "山水蒙", Symbol: "䷃", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 0, 0, 1}},
	"风水涣":  {Name: "风水涣", Symbol: "䷺", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 0, 1, 1}},
	"天水讼":  {Name: "天水讼", Symbol: "䷅", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 1, 1, 1}},
	"天火同人": {Name: "天火同人", Symbol: "䷌", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 1, 1, 1}},

	// 震宫八卦
	"震为雷":  {Name: "震为雷", Symbol: "䷲", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 0, 0, 1}},
	"雷地豫":  {Name: "雷地豫", Symbol: "䷏", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 0, 1}},
	"雷水解":  {Name: "雷水解", Symbol: "䷧", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 0, 0, 1}},
	"雷风恒":  {Name: "雷风恒", Symbol: "䷟", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 0, 0, 1}},
	"地风升":  {Name: "地风升", Symbol: "䷭", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 0, 0, 0}},
	"水风井":  {Name: "水风井", Symbol: "䷯", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 0, 1, 0}},
	"泽风大过": {Name: "泽风大过", Symbol: "䷛", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 1, 1, 0}},
	"泽雷随":  {Name: "泽雷随", Symbol: "䷐", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 1, 1, 0}},

	// 巽宫八卦
	"巽为风":  {Name: "巽为风", Symbol: "䷸", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 0, 1, 1}},
	"风天小畜": {Name: "风天小畜", Symbol: "䷈", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 0, 1, 1}},
	"风火家人": {Name: "风火家人", Symbol: "䷤", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 0, 1, 1}},
	"风雷益":  {Name: "风雷益", Symbol: "䷩", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 0, 1, 1}},
	"天雷无妄": {Name: "天雷无妄", Symbol: "䷘", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 1, 1, 1}},
	"火雷噬嗑": {Name: "火雷噬嗑", Symbol: "䷔", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 1, 0, 1}},
	"山雷颐":  {Name: "山雷颐", Symbol: "䷚", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 1, 0, 0}},
	"山风蛊":  {Name: "山风蛊", Symbol: "䷑", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[5], Yao: [6]int{0, 1, 1, 1, 0, 0}},

	// 坎宫八卦
	"坎为水":  {Name: "坎为水", Symbol: "䷜", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 0, 1, 0}},
	"水泽节":  {Name: "水泽节", Symbol: "䷻", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 0, 1, 0}},
	"水雷屯":  {Name: "水雷屯", Symbol: "䷂", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 0, 1, 0}},
	"水火既济": {Name: "水火既济", Symbol: "䷾", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 0, 1, 0}},
	"泽火革":  {Name: "泽火革", Symbol: "䷰", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 1, 1, 0}},
	"雷火丰":  {Name: "雷火丰", Symbol: "䷶", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 0, 0, 1}},
	"地火明夷": {Name: "地火明夷", Symbol: "䷣", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 0, 0, 0}},
	"地水师":  {Name: "地水师", Symbol: "䷆", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[6], Yao: [6]int{0, 1, 0, 0, 0, 0}},

	// 艮宫八卦
	"艮为山":  {Name: "艮为山", Symbol: "䷳", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[7], Yao: [6]int{1, 0, 0, 1, 0, 0}},
	"山火贲":  {Name: "山火贲", Symbol: "䷕", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[3], Yao: [6]int{1, 0, 1, 1, 0, 0}},
	"山天大畜": {Name: "山天大畜", Symbol: "䷙", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 1, 0, 0}},
	"山泽损":  {Name: "山泽损", Symbol: "䷨", UpperGua: BaGuaMap[7], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 1, 0, 0}},
	"火泽睽":  {Name: "火泽睽", Symbol: "䷥", UpperGua: BaGuaMap[3], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 1, 0, 1}},
	"天泽履":  {Name: "天泽履", Symbol: "䷉", UpperGua: BaGuaMap[1], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 1, 1, 1}},
	"风泽中孚": {Name: "风泽中孚", Symbol: "䷼", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 0, 1, 1}},
	"风山渐":  {Name: "风山渐", Symbol: "䷴", UpperGua: BaGuaMap[5], LowerGua: BaGuaMap[7], Yao: [6]int{1, 0, 0, 0, 1, 1}},

	// 坤宫八卦
	"坤为地":  {Name: "坤为地", Symbol: "䷁", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 0, 0}},
	"地雷复":  {Name: "地雷复", Symbol: "䷗", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[4], Yao: [6]int{0, 0, 1, 0, 0, 0}},
	"地泽临":  {Name: "地泽临", Symbol: "䷒", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[2], Yao: [6]int{1, 1, 0, 0, 0, 0}},
	"地天泰":  {Name: "地天泰", Symbol: "䷊", UpperGua: BaGuaMap[8], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 0, 0, 0}},
	"雷天大壮": {Name: "雷天大壮", Symbol: "䷡", UpperGua: BaGuaMap[4], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 0, 0, 1}},
	"泽天夬":  {Name: "泽天夬", Symbol: "䷪", UpperGua: BaGuaMap[2], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 1, 1, 0}},
	"水天需":  {Name: "水天需", Symbol: "䷄", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[1], Yao: [6]int{1, 1, 1, 0, 1, 0}},
	"水地比":  {Name: "水地比", Symbol: "䷇", UpperGua: BaGuaMap[6], LowerGua: BaGuaMap[8], Yao: [6]int{0, 0, 0, 0, 1, 0}},
}

// 计算先天卦
func calculateXianTianGua(bz BaZi) (*Hexagram, int) {
	yearNum := DiZhiToNumber[bz.YearZhi]
	monthNum := DiZhiToNumber[bz.MonthZhi]
	dayNum := DiZhiToNumber[bz.DayZhi]
	hourNum := DiZhiToNumber[bz.HourZhi]

	// 计算上卦
	upperSum := yearNum + monthNum + dayNum
	upperRemainder := upperSum % 8
	if upperRemainder == 0 {
		upperRemainder = 8
	}

	// 计算下卦
	lowerSum := yearNum + monthNum + dayNum + hourNum
	lowerRemainder := lowerSum % 8
	if lowerRemainder == 0 {
		lowerRemainder = 8
	}

	// 计算动爻
	yaoSum := yearNum + monthNum + dayNum + hourNum
	movingYao := yaoSum % 6
	if movingYao == 0 {
		movingYao = 6
	}

	// 获取上下卦
	upperGua := BaGuaMap[upperRemainder]
	lowerGua := BaGuaMap[lowerRemainder]

	// 构建卦名
	guaName := upperGua.Name + lowerGua.Name

	// 查找六十四卦
	if hexagram, exists := Gua64Map[guaName]; exists {
		return hexagram, movingYao
	}

	// 如果找不到，创建一个新的卦
	return &Hexagram{
		Name:     guaName,
		UpperGua: upperGua,
		LowerGua: lowerGua,
	}, movingYao
}

// 计算后天卦
func calculateHouTianGua(xianTianGua *Hexagram, movingYao int, bz BaZi) *Hexagram {
	// 复制先天卦的爻
	var newYao [6]int
	copy(newYao[:], xianTianGua.Yao[:])

	// 元堂爻阴阳互换
	if newYao[movingYao-1] == 0 {
		newYao[movingYao-1] = 1
	} else {
		newYao[movingYao-1] = 0
	}

	// 上下卦互换
	newUpperGua := findGuaByYao([3]int{newYao[3], newYao[4], newYao[5]})
	newLowerGua := findGuaByYao([3]int{newYao[0], newYao[1], newYao[2]})

	// 构建卦名
	guaName := newUpperGua.Name + newLowerGua.Name

	// 查找六十四卦
	if hexagram, exists := Gua64Map[guaName]; exists {
		return hexagram
	}

	// 如果找不到，创建一个新的卦
	return &Hexagram{
		Name:     guaName,
		UpperGua: newUpperGua,
		LowerGua: newLowerGua,
		Yao:      newYao,
	}
}

// 根据爻查找卦
func findGuaByYao(yao [3]int) *Gua {
	for _, gua := range BaGuaMap {
		if gua.Yao == yao {
			return gua
		}
	}
	return nil
}

// 计算流年卦（简化版）
func calculateLiuNianGua(xianTianGua *Hexagram, houTianGua *Hexagram, year int) string {
	// 在实际应用中，流年卦需要结合大运和具体年份计算
	// 这里只是一个简单的示例
	yearLastDigit := year % 10

	if yearLastDigit%2 == 0 {
		return fmt.Sprintf("流年卦以%s为主，需注意人际关系和健康", houTianGua.Name)
	} else {
		return fmt.Sprintf("流年卦以%s为主，利于事业发展和财运", xianTianGua.Name)
	}
}

// 主函数
func main() {
	// 示例八字：壬申年四月十一日巳时（男性）
	bz := BaZi{
		YearZhi:  "申",
		MonthZhi: "卯", // 四月
		DayZhi:   "辰", // 十一日假设为子日
		HourZhi:  "寅",
		Gender:   "男",
	}

	fmt.Println("输入八字信息:")
	fmt.Printf("年支: %s, 月支: %s, 日支: %s, 时支: %s, 性别: %s\n",
		bz.YearZhi, bz.MonthZhi, bz.DayZhi, bz.HourZhi, bz.Gender)

	// 计算先天卦
	xianTianGua, movingYao := calculateXianTianGua(bz)
	fmt.Printf("\n先天卦: %s%s, 动爻: %d爻\n", xianTianGua.UpperGua.Name, xianTianGua.LowerGua.Name, movingYao)

	// 计算后天卦
	houTianGua := calculateHouTianGua(xianTianGua, movingYao, bz)
	fmt.Printf("后天卦: %s%s\n", houTianGua.UpperGua.Name, houTianGua.LowerGua.Name)

	// 计算流年卦（以2023年为例）
	liuNian := calculateLiuNianGua(xianTianGua, houTianGua, 2023)
	fmt.Printf("流年卦: %s\n", liuNian)

	// 显示更多解释
	fmt.Println("\n卦象解释:")
	fmt.Printf("先天卦(%s): 代表命主前半生命运格局\n", xianTianGua.Name)
	fmt.Printf("后天卦(%s): 代表命主后半生命运变化\n", houTianGua.Name)
	fmt.Println("注：此为简化计算，实际应用需结合更多因素综合分析")
}

// 辅助函数：显示卦的详细信息
func displayGuaInfo(gua *Hexagram) string {
	var yaoStr []string
	for _, yao := range gua.Yao {
		if yao == 0 {
			yaoStr = append(yaoStr, "阴")
		} else {
			yaoStr = append(yaoStr, "阳")
		}
	}
	return fmt.Sprintf("%s(%s): %s", gua.Name, gua.Symbol, strings.Join(yaoStr, ""))
}
