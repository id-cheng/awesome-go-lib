package main

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func main() {
	// open an existing file
	wb, _ := xlsx.OpenFile("../samplefile.xlsx")
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
}
