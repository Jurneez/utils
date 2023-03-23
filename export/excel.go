package export

import (
	"fmt"
	"tool/utils/reflects"

	"github.com/tealeg/xlsx"
)

type Excel struct {
	Headers  []string
	Data     interface{} // 数组类型 结构体
	FilePath string      //文件保存路径
	Sheet    *xlsx.Sheet
	File     *xlsx.File
}

/*
path: 数据导出路径
*/
func NewExcel(headers []string, data interface{}, path string) (*Excel, error) {
	if path == "" {
		path = "./excel.xlsx"
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		return nil, err
	}

	return &Excel{
		Headers:  headers,
		Data:     data,
		File:     file,
		FilePath: path,
		Sheet:    sheet,
	}, nil
}

func (e *Excel) Export() error {
	dataMap, err := reflects.StructToMapSlice(e.Data, "")
	if err != nil {
		return err
	}
	if len(dataMap) == 0 {
		return nil
	}

	if len(e.Headers) == 0 {
		headers := make([]string, 0)
		for k := range dataMap[0] {
			headers = append(headers, k)
		}
		e.Headers = headers
	}

	headerRow := e.Sheet.AddRow() // 创建首行
	headerRow.SetHeightCM(0.5)    // 设置行高

	for _, each := range e.Headers {
		headerRow.AddCell().Value = each
	}

	for _, index := range dataMap {
		row := e.Sheet.AddRow()
		row.SetHeightCM(0.5)
		for _, key := range e.Headers {
			row.AddCell().Value = fmt.Sprintf("%v", index[key])
		}
	}

	// 持久化到磁盘
	if err := e.File.Save("data.xlsx"); err != nil {
		return err
	}

	return nil
}
