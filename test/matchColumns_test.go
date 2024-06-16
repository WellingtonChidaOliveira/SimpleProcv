package test

import (
	"github.com/wellingtonchida/simpleprocv/internal/data"
	"testing"
)

var (
	value             = "Female"
	templatePathFirst = "../templates/sample.xlsx"
	sheet             = "Sheet1"
)

func TestMatchColumns(t *testing.T) {

	t.Run("TestMatchColumns", func(t *testing.T) {
		got, _ := data.MatchColumns(value, templatePathFirst, sheet)
		want := []int{2, 3, 5, 6, 8, 9, 10}

		if len(got) != len(want) {
			t.Errorf("got %v \nwant %v ", got, want)
		}
	})

	t.Run("InvalidPath", func(t *testing.T) {
		_, err := data.MatchColumns(value, "invalidPath", sheet)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("InvalidSheet", func(t *testing.T) {
		_, err := data.MatchColumns(value, templatePathFirst, "invalidSheet")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
