package main

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func main() {
	// open an existing file
	wb, _ := xlsx.OpenFile("../samplefile.xlsx")
	sheetName := "Sample"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Println("Max row in sheet:", sh.MaxRow)
}
