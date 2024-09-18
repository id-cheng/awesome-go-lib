package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"strings"
)

var csv = `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`

func main() {
	df := dataframe.ReadCSV(strings.NewReader(csv))
	fmt.Println(df.String())
}
