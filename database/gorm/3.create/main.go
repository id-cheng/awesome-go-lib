package main

import (
	"fmt"
	"github.com/id-cheng/awesome-go-lib/database/gorm/2.model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func main() {
	// 创建记录
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	user := model.User{Name: "Cheng", Age: 18, Birthday: time.Now()}
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

	users := []*model.User{
		{Name: "Cheng", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	db.Create(users)

	//用指定的字段创建记录
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	//创建记录并忽略传递给 ‘Omit’ 的字段值
	db.Omit("Name", "Age", "CreatedAt").Create(&user)
	//批量插入
	db.CreateInBatches(users, 100)

	//使用CreateBatchSize 选项初始化GORM实例后，此后进行创建 & 关联操作时所有的INSERT行为都会遵循初始化时的配置
	db, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		CreateBatchSize: 1000,
	})
	db = db.Session(&gorm.Session{CreateBatchSize: 1000})

	//支持通过 map[string]any 和 []map[string]any{}来创建记录
	db.Model(&model.User{}).Create(map[string]any{
		"Name": "cheng", "Age": 18,
	})

	// batch insert from `[]map[string]any{}`
	db.Model(&model.User{}).Create([]map[string]any{
		{"Name": "cheng1", "Age": 18},
		{"Name": "cheng2", "Age": 20},
	})

	//创建关联数据时，如果关联值非零，这些关联会被upsert，并且它们的Hooks方法也会被调用。
	type CreditCard struct {
		gorm.Model
		Number string
		UserID uint
	}

	type User struct {
		gorm.Model
		Name       string
		CreditCard CreditCard
	}

	db.Create(&User{
		Name:       "cheng",
		CreditCard: CreditCard{Number: "123"},
	})
	// INSERT INTO `users` ...
	// INSERT INTO `credit_cards` ...

	// 通过Select, Omit方法来跳过关联更新
	db.Omit("CreditCard").Create(&user)

	// skip all associations
	db.Omit(clause.Associations).Create(&user)

}
