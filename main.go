package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Let's Start")
	// gorm.Open(mysql.Open("データベースの使用許可のあるUser:mysqlに入る際のパスワード@/go_basics")
	db, err := gorm.Open(mysql.Open("root:RootMyroot7@/go_basics"), &gorm.Config{})

	if err != nil {
		panic("データベースに接続できませんでした。")
	}

	fmt.Println(db)

	db.AutoMigrate(&User{})

	user3 := User{
		Id: 3,
	}
	db.Delete(&user3)
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
