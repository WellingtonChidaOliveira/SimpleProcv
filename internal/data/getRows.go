package data

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func GetValuesInRows(getColumn, path, sheet string, rows []int) ([]string, error) {
	columnsInRows, err := getColumnValues(path, sheet, getColumn, rows)
	if err != nil {
		return nil, err
	}
	return columnsInRows, nil
}

func getColumnValues(path, sheet, column string, rows []int) ([]string, error) {
	f, err := openFile(path)
	if err != nil {
		return nil, err
	}

	var columnValues []string
	for _, v := range rows {
		columnNumber := fmt.Sprintf("%v%d", column, v)

		cell, err := f.GetCellValue(sheet, columnNumber)
		if err != nil {
			return nil, err
		}
		columnValues = append(columnValues, cell)
	}

	return columnValues, nil
}

func openFile(path string) (*excelize.File, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
