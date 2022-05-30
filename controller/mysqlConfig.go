package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Favorite struct {
	Id      int64  `json:"id,omitempty"`
	UserId  int64  `gorm:"not null"`
	VideoId string `gorm:"not null"`
}

var Db *gorm.DB

func InitDb() {
	Db = createDb()

}

func createDb() *gorm.DB {
	dsn := "root:dir99@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("失败")
	}

	return db
}
