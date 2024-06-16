package data

import (
	"fmt"
	"log/slog"
)

func SetValuesInRows(setColumn, path, sheet string, rows []int, values []string) error {
	f, err := openFile(path)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	for i, row := range rows {
		columnNumber := fmt.Sprintf("%v%d", setColumn, row)
		err := f.SetCellValue(sheet, columnNumber, values[i])
		if err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	err = f.Save()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil

}
