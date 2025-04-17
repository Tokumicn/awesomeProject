package main

import (
	"github.com/xuri/excelize/v2"
	"math/rand"
)

func main() {
	f := excelize.NewFile()
	defer f.Close()

	// 预定义样式
	headerStyle := createHeaderStyle(f)
	rowStyle := createRowStyle(f)

	sw, _ := f.NewStreamWriter("Sheet1")
	defer sw.Flush()

	// 写入标题
	sw.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: headerStyle, Value: "ID"},
		excelize.Cell{StyleID: headerStyle, Value: "销售额"},
	})

	// 批量写入数据
	for row := 2; row <= 100000; row++ {
		cell, _ := excelize.CoordinatesToCellName(1, row)
		sw.SetRow(cell, []interface{}{
			row - 1,
			rand.Float64() * 10000,
		}, excelize.RowOpts{StyleID: rowStyle})
	}

	f.SaveAs("sales_report.xlsx")
}

func createHeaderStyle(f *excelize.File) int {
	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#2F5496"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Border: []excelize.Border{
			{Type: "bottom", Color: "#000000", Style: 2},
		},
	})
	return style
}

func createRowStyle(f *excelize.File) int {
	rowStyle, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "bottom", Color: "#D3D3D3", Style: 1},
		},
		NumFmt: 4, // 货币格式
	})

	return rowStyle
}
