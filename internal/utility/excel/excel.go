package excel

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xuri/excelize/v2"
)

var (
	ctx = gctx.GetInitCtx()
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
	Sw         *excelize.StreamWriter
	sheetName  string
}

// 实例化数据
func NewExcel() *Excel {
	newExcel := excelize.NewFile()
	return &Excel{File: newExcel}
}

// 设置sheet名
func (e *Excel) SetSheet(sheet string) *Excel {
	e.sheetName = sheet
	// 创建一个工作表
	_, err := e.File.NewSheet(sheet)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	return e
}

func (e *Excel) SetSheetSteam(sheet string) *Excel {
	sw, err := e.File.NewStreamWriter(sheet)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	e.Sw = sw
	e.sheetName = sheet
	return e
}

// 设置样式
func (e *Excel) SetStyle(len int) {
	columnName, err := excelize.ColumnNumberToName(len)
	if err != nil {
		g.Log().Error(ctx, err)
		return
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
		return
	}

	//设置加粗
	err = e.File.SetRowStyle(e.sheetName, 1, 1, styleHeadInt)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	err = e.File.SetColWidth(e.sheetName, "A", columnName, 18)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
}

// 设置表单表头
func (e *Excel) SetCellHead(dataHead []interface{}) *Excel {
	if e.Sw == nil {
		for i, v := range dataHead {
			columnName, err := excelize.ColumnNumberToName(i + 1)
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
			err = e.File.SetCellValue(e.sheetName, fmt.Sprintf("%v%v", columnName, 1), v)
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
		}
		//设置样式
		e.SetStyle(len(dataHead))
	} else {
		styleID, err := e.File.NewStyle(&excelize.Style{
			Font: &excelize.Font{
				Bold: true,
				Size: float64(14),
			},
		})
		if err != nil {
			fmt.Println(err)
			g.Log().Error(ctx, err)
			return nil
		}
		var list []interface{}
		for _, v := range dataHead {
			cell := excelize.Cell{StyleID: styleID, Value: v}
			list = append(list, cell)
		}
		err = e.Sw.SetRow("A1", list)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}
	}

	return e
}

// 设置表单数据 同步写入
func (e *Excel) SetCellRow(data [][]interface{}, startNum int) *Excel {

	if startNum < 0 {
		startNum = 0
	}
	rowIndex := startNum
	for _, rowValue := range data {
		rowIndex++
		if e.Sw == nil {
			for i, v := range rowValue {
				columnName, err := excelize.ColumnNumberToName(i + 1)
				if err != nil {
					g.Log().Error(ctx, err)
					return nil
				}
				err = e.File.SetCellValue(e.sheetName, fmt.Sprintf("%v%v", columnName, rowIndex), v)
				if err != nil {
					g.Log().Error(ctx, err)
					return nil
				}
			}
		} else {
			cell, err := excelize.CoordinatesToCellName(1, rowIndex)
			if err != nil {
				fmt.Println(err)
				break
			}
			err = e.Sw.SetRow(cell, rowValue)
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}
	}

	g.Log().Info(ctx, "Wrote rows from", startNum, "to", rowIndex-1)
	return e
}

func (e *Excel) SetCellRowsBatch(data [][]interface{}, startNum int, batchSize int) *Excel {
	if batchSize == 0 {
		e.SetCellRow(data, startNum)
		err := e.File.Save()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}
	} else {
		totalRows := len(data)
		for startNum = 0; startNum < totalRows; startNum += batchSize {
			end := startNum + batchSize
			if end > totalRows {
				end = totalRows
			}
			g.Log().Debug(ctx, "Processing batch from", startNum+1, "to", end)
			e.SetCellRow(data[startNum:end], startNum)
			err := e.File.Save()
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
		}
	}

	return e
}

// 设置批量表格数据
func (e *Excel) SetExcelSheetData(excelValue []*ExcelSheet) *Excel {
	e.ExcelSheet = excelValue
	for _, v := range e.ExcelSheet {
		e.SetSheet(v.SheetName)
		e.SetCellHead(v.SheetHead)
		e.SetCellRow(v.SheetList, 1)
	}
	return e
}

// 生成Excel文件
func (e *Excel) SaveExcel(path string, fileName string) (pathFileName string) {

	if e.Sw != nil {
		if err := e.Sw.Flush(); err != nil {
			g.Log().Error(ctx, err)
			return
		}
		e.Sw = nil
		g.Log().Info(ctx, "Flush 结束")
	}

	pathFileName = fmt.Sprintf("%v%v.xlsx", path, fileName)
	// 根据指定路径保存文件
	if err := e.File.SaveAs(pathFileName); err != nil {
		g.Log().Error(ctx, err)
	}

	// 关闭文件释放资源
	if err := e.File.Close(); err != nil {
		g.Log().Error(ctx, err)
	}

	return "/" + pathFileName

}
