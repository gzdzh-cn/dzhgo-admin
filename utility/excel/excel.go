package excel

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xuri/excelize/v2"
)

var (
	ctx       = gctx.GetInitCtx()
	sheetName string
)

// 表sheet数据
type ExcelSheet struct {
	SheetName string
	SheetHead []interface{}
	SheetList [][]interface{}
}

type Excel struct {
	File       *excelize.File
	ExcelSheet []*ExcelSheet
}

// 实例化数据
func NewExcel() *Excel {
	newExcel := excelize.NewFile()
	return &Excel{File: newExcel}
}

// 设置sheet名
func (e *Excel) SetSheet(sheet string) *Excel {
	sheetName = sheet
	// 创建一个工作表
	_, err := e.File.NewSheet(sheet)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return e
}

// 设置样式
func (e *Excel) SetStyle(len int) {
	columnName, err := excelize.ColumnNumberToName(len)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	headStyle := &excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: float64(14),
		},
	}

	styleHeadInt, err := e.File.NewStyle(headStyle)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	//设置加粗
	err = e.File.SetRowStyle(sheetName, 1, 1, styleHeadInt)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	e.File.SetColWidth(sheetName, "A", columnName, 18)
}

// 设置表单表头
func (e *Excel) SetCellHead(dataHead []interface{}) *Excel {

	for i, v := range dataHead {
		columnName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}
		err = e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v", columnName, 1), v)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}
	}

	//设置样式
	e.SetStyle(len(dataHead))

	return e
}

// 设置表单数据
func (e *Excel) SetCellRow(data [][]interface{}) *Excel {

	rowIndex := 1
	for _, rowValue := range data {
		rowIndex++
		for i, v := range rowValue {
			columnName, err := excelize.ColumnNumberToName(i + 1)
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
			err = e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v", columnName, rowIndex), v)
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
		}
	}
	return e
}

// 生成Excel文件
func (e *Excel) SaveExcel(path string) (pathFileName string) {
	time := gtime.Now()
	pathFileName = fmt.Sprintf("%v%v.xlsx", path, time.Format("YmdHisu"))
	// 根据指定路径保存文件
	if err := e.File.SaveAs(pathFileName); err != nil {
		g.Log().Error(ctx, err)
	}

	return "/" + pathFileName

}

// 设置批量表格数据
func (e *Excel) SetExcelSheetData(excelValue []*ExcelSheet) *Excel {
	e.ExcelSheet = excelValue
	for _, v := range e.ExcelSheet {
		e.File.NewSheet(v.SheetName)
		sheetName = v.SheetName
		e.SetCellHead(v.SheetHead)
		e.SetCellRow(v.SheetList)
	}
	return e
}
