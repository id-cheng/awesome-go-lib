package main

import (
	"fmt"

	"github.com/cch123/elasticsql"
)

var sql = `select * from project
		   where id = 1 and name = 'cheng'
		   and create_time between '2020-01-01T00:00:00+0800' 
		   and '2022-01-01T00:00:00+0800'
		   and category_id > 1 
		   order by id desc limit 100,10
`

func main() {
	dsl, table, _ := elasticsql.Convert(sql)
	fmt.Println("dsl:", dsl)
	fmt.Println("table:", table)
}
