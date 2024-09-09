package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer f.Close()
	index, _ := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book.xlsx"); err != nil {
		fmt.Println(err)
	}
}
