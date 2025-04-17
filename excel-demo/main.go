package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer f.Close()

	// 定义标题样式
	styleID, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 14, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#4F81BD"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	// 写入标题行并应用样式
	headers := []string{"姓名", "年龄", "部门"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue("Sheet1", cell, h)
	}
	f.SetCellStyle("Sheet1", "A1", "C1", styleID)

	// 保存文件
	if err := f.SaveAs("员工信息.xlsx"); err != nil {
		fmt.Println(err)
	}
}
