package main

import (
	"errors"
	"fmt"
	"github.com/id-cheng/awesome-go-lib/database/gorm/2.model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 获取第一条记录（主键升序）
	db.First(&model.User{})
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&model.User{})
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&model.User{})
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&model.User{})

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 根据主键检索
	db.First(&model.User{}, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&model.User{}, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&model.User{}, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	db.First(&model.User{}, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")

	var user = model.User{ID: 10}
	db.First(&user)
	// SELECT * FROM users WHERE id = 10;

	var res model.User
	db.Model(model.User{ID: 10}).First(&res)
	// SELECT * FROM users WHERE id = 10;
	fmt.Println(res)
}
