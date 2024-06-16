package main

import (
	"fmt"

	"github.com/wellingtonchida/simpleprocv/internal/data"
)

func main() {
	fmt.Println("Hello, World!")

	var indexColumns, mainFilePath, secondFilePath, sheetName, columnToGetValue, columnToSetValue string

	fmt.Print("Enter index columns: ")
	fmt.Scanln(&indexColumns)

	fmt.Print("Enter main file path: ")
	fmt.Scanln(&mainFilePath)

	fmt.Print("Enter second file path: ")
	fmt.Scanln(&secondFilePath)

	fmt.Print("Enter sheet name: ")
	fmt.Scanln(&sheetName)

	fmt.Print("Enter column to get value: ")
	fmt.Scanln(&columnToGetValue)

	fmt.Print("Enter column to set value: ")
	fmt.Scanln(&columnToSetValue)

	rowsMateched, _ := data.MatchColumns(indexColumns, mainFilePath, sheetName)
	valuesToSet, _ := data.GetValuesInRows(columnToGetValue, mainFilePath, sheetName, rowsMateched)
	err := data.SetValuesInRows(columnToSetValue, secondFilePath, sheetName, rowsMateched, valuesToSet)
	if err != nil {
		return
	}

	fmt.Println("Done")
}

// var (
// 	indexColumns     = "Female"
// 	mainFilePath     = "../../templates/sample.xlsx"
// 	secondFilePath   = "../../templates/sample2.xlsx"
// 	sheetName        = "Sheet1"
// 	columnToGetValue = "A"
// 	columnToSetValue = "A"
// )
