package test

import (
	"testing"

	"github.com/wellingtonchida/simpleprocv/internal/data"
)

var (
	//indexColumns     = "Female"
	mainFilePath = "../templates/sample.xlsx"
	//secondFilePath   = "../templates/sample2.xlsx"
	sheetName        = "Sheet1"
	columnToGetValue = "A"
	//columnToSetValue = "A"
)

func TestGetExcel(t *testing.T) {

	t.Run("GetValuesInRows", func(t *testing.T) {
		got, _ := data.GetValuesInRows(columnToGetValue, mainFilePath, sheetName, []int{1, 2, 3})
		//want := []string{}

		if len(got) <= 0 {
			t.Errorf("got %v want %v", got, "err")
		}
	})

	t.Run("InvalidPath", func(t *testing.T) {
		_, got := data.GetValuesInRows(columnToGetValue, "invalidPath", sheetName, []int{1, 2, 3})
		want := "open invalidPath: The system cannot find the file specified."

		if got.Error() != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("InvalidSheetName", func(t *testing.T) {
		_, got := data.GetValuesInRows(columnToGetValue, mainFilePath, "Sheet2", []int{1, 2, 3})
		want := "sheet Sheet2 does not exist"
		if got.Error() != want {
			t.Errorf("got %v want %v", got, want)

		}
	})
	t.Run("InvalidColumn", func(t *testing.T) {
		_, got := data.GetValuesInRows("P", mainFilePath, sheetName, []int{1, 2, 3})
		//want := "column 1 does not exist"
		if got != nil {
			t.Errorf("Got should be get error")
		}
	})
}
