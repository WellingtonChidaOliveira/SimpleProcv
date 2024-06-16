package test

import (
	"github.com/wellingtonchida/simpleprocv/internal/data"
	"testing"
)

func TestSetValuesInRows(t *testing.T) {
	tests := []struct {
		name          string
		setColumn     string
		path          string
		sheet         string
		rows          []int
		values        []string
		expectedError string
	}{
		{
			name:          "ValidCase",
			setColumn:     "A",
			path:          "../templates/sample.xlsx",
			sheet:         "Sheet1",
			rows:          []int{1, 2, 3},
			values:        []string{"value1", "value2", "value3"},
			expectedError: "",
		},
		{
			name:          "InvalidPath",
			setColumn:     "A",
			path:          "invalidPath",
			sheet:         "Sheet1",
			rows:          []int{1, 2, 3},
			values:        []string{"value1", "value2", "value3"},
			expectedError: "open invalidPath: The system cannot find the file specified.",
		},
		{
			name:          "InvalidSheetName",
			setColumn:     "A",
			path:          "../templates/sample.xlsx",
			sheet:         "Sheet2",
			rows:          []int{1, 2, 3},
			values:        []string{"value1", "value2", "value3"},
			expectedError: "sheet Sheet2 does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := data.SetValuesInRows(tt.setColumn, tt.path, tt.sheet, tt.rows, tt.values)
			if err != nil {
				if err.Error() != tt.expectedError {
					t.Errorf("got error %v, want %v", err, tt.expectedError)
				}
				return
			}
			if tt.expectedError != "" {
				t.Errorf("expected error %v, got nil", tt.expectedError)
			}
		})
	}
}
