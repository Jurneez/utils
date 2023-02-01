package excel

import (
	"errors"

	"github.com/tealeg/xlsx"
)

type Excel struct {
	Sheets   []*Sheet
	FilePath string
	File     *xlsx.File
}
type Sheet struct {
	name   string
	titles []string // excel的头部名称
	cells  [][]interface{}
}

func NewExcel() *Excel {
	return &Excel{
		Sheets:   make([]*Sheet, 0),
		FilePath: "./excel.xlsx",
		File:     xlsx.NewFile(),
	}
}

func (e *Excel) SetFilePath(path string) {
	e.FilePath = path
}

func (e *Excel) SetSheet(sheetName string) error {
	sheet, err := e.File.AddSheet(sheetName)
	if err != nil {
		return err
	}
	if sheet == nil {
		return errors.New("error")
	}

	return nil
}

func Excel1() {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "xixi"

	err := file.Save("./file.xlsx")
	if err != nil {
		panic(err)
	}

}
