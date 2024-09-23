package main

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func main() {
	filename := "samplefile.xlsx"
	wb, _ := xlsx.OpenFile(filename)
	sh, _ := wb.AddSheet("My New Sheet")
	sh, _ = wb.AppendSheet(*sh, "A new sheet")
	fmt.Println(sh)
}
