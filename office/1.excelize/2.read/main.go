package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, _ := excelize.OpenFile("Book.xlsx")
	defer f.Close()
	// Get value from cell by given worksheet name and cell reference.
	cell, _ := f.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, _ := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
