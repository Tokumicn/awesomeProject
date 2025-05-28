package main

import (
	"ai-demo/OCR/ocr-demo1/file_utils"
	"ai-demo/OCR/ocr-demo1/ocr_cli"
	strings_pipeline "ai-demo/OCR/ocr-demo1/string-data-pipeline"
	"context"
	"fmt"
	"log/slog"
	"os"
)

func init() {
	l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(l)
}

var rawJson = `[{"text":"汇灵盏","score":0.9666414260864258,"position":[[83,18],[150,18],[150,43],[83,43]]},{"text":"五色旗盒","score":0.8097074627876282,"position":[[314,19],[400,19],[400,42],[314,42]]},{"text":"碧玉葫芦","score":0.9930232763290405,"position":[[541,17],[629,16],[630,42],[542,44]]},{"text":"碧玉葫芦","score":0.8266566395759583,"position":[[771,19],[858,16],[859,42],[772,44]]},{"text":"单价","score":0.9177404046058655,"position":[[82,58],[121,58],[121,80],[82,80]]},{"text":"1599999","score":0.8373721837997437,"position":[[133,61],[196,61],[196,78],[133,78]]},{"text":"单价","score":0.9924006462097168,"position":[[311,58],[351,58],[351,80],[311,80]]},{"text":"1599999","score":0.5702419877052307,"position":[[361,61],[425,61],[425,78],[361,78]]},{"text":"单价","score":0.9831734299659729,"position":[[541,58],[581,58],[581,80],[541,80]]},{"text":"1599999","score":0.9715652465820312,"position":[[592,61],[656,61],[656,78],[592,78]]},{"text":"单价","score":0.9919344186782837,"position":[[770,58],[811,58],[811,80],[770,80]]},{"text":"1599999","score":0.5073329210281372,"position":[[822,61],[884,61],[884,78],[822,78]]},{"text":"汇灵盏","score":0.9858004450798035,"position":[[83,109],[150,111],[149,137],[82,135]]},{"text":"金甲仙衣","score":0.7190647721290588,"position":[[313,112],[400,112],[400,136],[313,136]]},{"text":"金甲仙衣","score":0.9508416056632996,"position":[[543,112],[630,112],[630,135],[543,135]]},{"text":"嗜血幡","score":0.9874117374420166,"position":[[772,112],[837,111],[838,134],[773,135]]},{"text":"单价","score":0.983636200428009,"position":[[82,150],[122,152],[121,173],[81,171]]},{"text":"1599999","score":0.9870477318763733,"position":[[133,154],[196,154],[196,172],[133,172]]},{"text":"单价","score":0.6670665740966797,"position":[[311,150],[351,150],[351,174],[311,174]]},{"text":"1799999","score":0.9119083285331726,"position":[[361,154],[425,154],[425,171],[361,171]]},{"text":"单价","score":0.9771367907524109,"position":[[541,150],[581,152],[580,173],[540,172]]},{"text":"1799999","score":0.8913893103599548,"position":[[591,154],[655,154],[655,171],[591,171]]},{"text":"单价","score":0.5605122447013855,"position":[[770,150],[811,150],[811,174],[770,174]]},{"text":"1899999","score":0.9539927244186401,"position":[[822,154],[885,154],[885,172],[822,172]]},{"text":"嗜血幡","score":0.9681030511856079,"position":[[83,205],[150,205],[150,229],[83,229]]},{"text":"拭剑石","score":0.8687690496444702,"position":[[313,205],[379,205],[379,230],[313,230]]},{"text":"拭剑石","score":0.9727485775947571,"position":[[541,205],[609,205],[609,230],[541,230]]},{"text":"飞剑","score":0.9370542168617249,"position":[[778,203],[819,203],[819,230],[778,230]]},{"text":"单价","score":0.583051860332489,"position":[[81,241],[122,244],[121,267],[80,265]]},{"text":"1899999","score":0.7931272387504578,"position":[[134,247],[196,247],[196,265],[134,265]]},{"text":"单价","score":0.5242389440536499,"position":[[311,241],[352,244],[350,267],[310,265]]},{"text":"1899999","score":0.9923750758171082,"position":[[362,247],[425,247],[425,265],[362,265]]},{"text":"单价","score":0.8695617914199829,"position":[[541,242],[582,244],[581,268],[540,266]]},{"text":"1899999","score":0.992048442363739,"position":[[592,247],[655,247],[655,265],[592,265]]},{"text":"单价","score":0.9921098351478577,"position":[[770,244],[810,244],[810,266],[770,266]]},{"text":"1899999","score":0.9308920502662659,"position":[[822,247],[884,247],[884,265],[822,265]]},{"text":"九黎战鼓","score":0.8435644507408142,"position":[[84,298],[171,298],[171,322],[84,322]]},{"text":"九黎战鼓","score":0.8127223253250122,"position":[[313,298],[400,298],[400,323],[313,323]]},{"text":"风袋","score":0.9707527160644531,"position":[[542,297],[589,297],[589,323],[542,323]]},{"text":"风袋","score":0.9950265288352966,"position":[[772,297],[819,297],[819,323],[772,323]]},{"text":"单价","score":0.9687031507492065,"position":[[82,336],[122,338],[121,359],[81,358]]},{"text":"2199999","score":0.6322600841522217,"position":[[131,340],[196,340],[196,358],[131,358]]},{"text":"单价","score":0.9952405691146851,"position":[[311,336],[350,338],[349,359],[310,358]]},{"text":"2199999","score":0.5728441476821899,"position":[[361,340],[425,340],[425,358],[361,358]]},{"text":"单价","score":0.8740032315254211,"position":[[542,336],[581,338],[580,360],[541,358]]},{"text":"2399999","score":0.8286696672439575,"position":[[590,340],[655,340],[655,358],[590,358]]},{"text":"单价","score":0.8820915818214417,"position":[[771,336],[810,338],[809,359],[770,358]]},{"text":"2399999","score":0.6017298102378845,"position":[[820,340],[884,340],[884,358],[820,358]]},{"text":"异域风情","score":0.9496907591819763,"position":[[84,391],[170,391],[170,414],[84,414]]},{"text":"风袋","score":0.9950049519538879,"position":[[313,390],[359,390],[359,416],[313,416]]},{"text":"鬼谷子","score":0.9930417537689209,"position":[[543,391],[608,391],[608,416],[543,416]]},{"text":"法宝任务书","score":0.6326383948326111,"position":[[774,392],[879,392],[879,414],[774,414]]},{"text":"单价","score":0.8968239426612854,"position":[[82,430],[121,430],[121,452],[82,452]]},{"text":"2399999","score":0.8627421259880066,"position":[[131,433],[196,433],[196,451],[131,451]]},{"text":"单价","score":0.9925353527069092,"position":[[311,430],[350,430],[350,452],[311,452]]},{"text":"2399999","score":0.7649140954017639,"position":[[361,433],[425,433],[425,451],[361,451]]},{"text":"单价","score":0.9214774370193481,"position":[[541,430],[580,430],[580,452],[541,452]]},{"text":"2999999","score":0.5285037159919739,"position":[[590,433],[655,433],[655,451],[590,451]]},{"text":"单价","score":0.9917512536048889,"position":[[770,430],[810,430],[810,452],[770,452]]},{"text":"2999999","score":0.6828131675720215,"position":[[820,433],[884,433],[884,451],[820,451]]}]`

func main() {
	//dir, _ := os.Getwd()
	//fmt.Println("[My-OCR] shell work dir: ", dir)

	ocr_cli.SetThresholds(50, 40)

	ctx := context.TODO()

	//exePath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//binDir := filepath.Dir(exePath)
	//
	//fmt.Println("[My-OCR] bin work dir: ", binDir)
	//dirPath := fmt.Sprintf("%s/images/", binDir)
	// 读取文件

	var (
		file       *os.File
		err        error
		dirEntries []os.DirEntry
		ocrRes     []ocr_cli.OCRResult
		table      [][]ocr_cli.TableCell
	)

	imageDir := "../images/"
	dirEntries, err = os.ReadDir(imageDir)
	if err != nil {
		slog.ErrorContext(ctx, "read dir err: ", err.Error())
		return
	}

	// 最终待写入文件结果
	writeArr := []string{}
	for _, entry := range dirEntries {

		// 文件夹跳过不处理
		if entry.IsDir() {
			continue
		}

		// 检查文件后缀名是图片
		if !file_utils.IsImageExt(entry) {
			continue
		}

		// 将相对路径转换为绝对路径
		fileName := entry.Name()
		fullFileName := file_utils.ConvRelative2FullPath(imageDir, fileName)

		// 打开文件
		file, err = os.Open(fullFileName)
		if err != nil {
			slog.ErrorContext(ctx, "open file err: ", err.Error())
			return
		}

		// 发送OCR请求
		ocrRes, err = ocr_cli.PostOCR(ctx, fileName, file)
		if err != nil {
			slog.ErrorContext(ctx, "post ocr err: ", err.Error())
			return
		}

		//ocrResBytes, err := json.Marshal(ocrRes)
		//if err != nil {
		//	slog.ErrorContext(ctx, "json marshal err: ", err.Error())
		//}

		//fmt.Println("===================================================")
		//fmt.Println(string(ocrResBytes))
		//fmt.Println("================================================")

		//err = json.Unmarshal([]byte(rawJson), &ocrRes)
		//if err != nil {
		//	slog.ErrorContext(ctx, "json unmarshal err: ", err.Error())
		//	return
		//}

		// 解析OCR结果
		table, err = ocr_cli.ParseOCRToTableWithFilter(ocrRes, map[string]struct{}{
			"单价": {},
		}, []string{"等级"})

		//table, err = ocr_cli.ParseOCRToTable(ocrRes)
		//if err != nil {
		//	slog.ErrorContext(ctx, "parse ocr err: ", err.Error())
		//	return
		//}

		fmt.Println("===================================================")
		ocr_cli.PrintOCRTable(table)
		fmt.Println("===================================================")
		// 将结果映射到实体

		for _, row := range table {
			for _, cell := range row {
				writeArr = append(writeArr, cell.Text)
			}
		}
	}

	// 关闭文件
	if file != nil {
		err = file.Close()
		if err != nil {
			slog.ErrorContext(ctx, "close file err: ", err.Error())
		}
	}

	// 输出实体
	err = file_utils.WriteLinesToFile(writeArr, "output.txt")
	if err != nil {
	}

	// 初始化流水线
	pipeline := &strings_pipeline.Pipeline{}
	pipeline.AddStep(strings_pipeline.Deduplicate)
	pipeline.AddStep(strings_pipeline.NormalizeNames)

	// 解析输入数据
	items := strings_pipeline.ParseInput(writeArr)

	// 运行流水线处理
	result := pipeline.Run(items)

	// 输出结果
	for _, product := range result {
		fmt.Printf("{ Name: %q, Prices: %v }\n", product.Name, product.Prices)
	}
}
