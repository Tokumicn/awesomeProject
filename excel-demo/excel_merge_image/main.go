package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 定义股票详细信息结构体
type StockDetail struct {
	Industry  string // 行业信息
	StockName string // 股票名
	StockCode string // 股票代码
	Reason    string // 入选理由
}

// 定义数据结构
type StockInfo struct {
	Number    int           // 序号
	ImagePath string        // 图片路径
	ImageName string        // 图片名
	Details   []StockDetail // 股票详细信息数组
}

// 准备数据
func prepareStockData(files []os.DirEntry) ([]StockInfo, error) {
	var stockInfos []StockInfo
	index := 1

	for _, file := range files {
		if !file.IsDir() && isImageFile(file.Name()) {
			imagePath := filepath.Join("images", file.Name())

			// 创建股票信息
			stockInfo := StockInfo{
				Number:    index,
				ImagePath: imagePath,
				ImageName: file.Name(),
				// 示例多行数据
				Details: []StockDetail{
					{
						Industry:  "科技行业",
						StockName: "示例股票1",
						StockCode: "000001",
						Reason:    "这是一支具有良好发展前景的股票",
					},
					{
						Industry:  "互联网",
						StockName: "示例股票2",
						StockCode: "000002",
						Reason:    "公司业绩持续增长",
					},
					{
						Industry:  "人工智能",
						StockName: "示例股票3",
						StockCode: "000003",
						Reason:    "行业龙头地位稳固",
					},
				},
			}

			stockInfos = append(stockInfos, stockInfo)
			index++
		}
	}

	return stockInfos, nil
}

// 写入Excel文件
func writeToExcel(stockInfos []StockInfo) error {
	// 创建新的Excel文件
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 设置列宽
	f.SetColWidth("Sheet1", "A", "A", 10) // 序号
	f.SetColWidth("Sheet1", "B", "B", 30) // 图片
	f.SetColWidth("Sheet1", "C", "C", 30) // 图片名
	f.SetColWidth("Sheet1", "D", "D", 20) // 行业信息
	f.SetColWidth("Sheet1", "E", "E", 20) // 股票名
	f.SetColWidth("Sheet1", "F", "F", 15) // 股票代码
	f.SetColWidth("Sheet1", "G", "G", 50) // 入选理由

	// 设置表头
	headers := []string{"序号", "图片", "图片名", "行业信息", "股票名", "股票代码", "入选理由"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue("Sheet1", cell, header)
	}

	row := 2 // 从第二行开始写入数据
	for _, info := range stockInfos {
		// 读取图片以获取尺寸
		img, err := readImage(info.ImagePath)
		if err != nil {
			fmt.Printf("读取图片失败 %s: %v\n", info.ImageName, err)
			continue
		}

		// 获取图片尺寸
		bounds := img.Bounds()
		height := bounds.Dy()

		// 插入图片
		cell := fmt.Sprintf("B%d", row)
		if err := f.AddPicture("Sheet1", cell, info.ImagePath, &excelize.GraphicOptions{
			AutoFit: true,
			ScaleX:  0.5,
			ScaleY:  0.5,
		}); err != nil {
			fmt.Printf("插入图片失败 %s: %v\n", info.ImageName, err)
			continue
		}

		// 设置行高以适应图片
		heightInPoints := float64(height) * 0.75 / 96.0 // 将像素转换为Excel点
		f.SetRowHeight("Sheet1", row, heightInPoints)

		// 写入其他信息
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), info.Number)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), info.ImageName)

		// 写入第一行详细信息
		if len(info.Details) > 0 {
			detail := info.Details[0]
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), detail.Industry)
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), detail.StockName)
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), detail.StockCode)
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), detail.Reason)
		}

		// 写入剩余行详细信息
		for i := 1; i < len(info.Details); i++ {
			row++
			detail := info.Details[i]
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), detail.Industry)
			f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), detail.StockName)
			f.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), detail.StockCode)
			f.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), detail.Reason)
		}

		// 合并单元格
		if len(info.Details) > 1 {
			// 合并序号、图片和图片名
			f.MergeCell("Sheet1", fmt.Sprintf("A%d", row-len(info.Details)+1), fmt.Sprintf("A%d", row))
			f.MergeCell("Sheet1", fmt.Sprintf("B%d", row-len(info.Details)+1), fmt.Sprintf("B%d", row))
			f.MergeCell("Sheet1", fmt.Sprintf("C%d", row-len(info.Details)+1), fmt.Sprintf("C%d", row))
		}

		row++
	}

	// 保存文件
	if err := f.SaveAs("stock_info.xlsx"); err != nil {
		return fmt.Errorf("保存Excel文件失败: %v", err)
	}

	return nil
}

func main() {
	// 读取images目录下的所有图片
	files, err := os.ReadDir("images")
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
		return
	}

	// 准备数据
	stockInfos, err := prepareStockData(files)
	if err != nil {
		fmt.Printf("准备数据失败: %v\n", err)
		return
	}

	// 写入Excel文件
	if err := writeToExcel(stockInfos); err != nil {
		fmt.Printf("生成Excel文件失败: %v\n", err)
		return
	}

	fmt.Println("Excel文件已生成: stock_info.xlsx")
}

// 检查文件是否为图片
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// 读取图片文件
func readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Decode(file)
	case ".png":
		return png.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", ext)
	}
}
