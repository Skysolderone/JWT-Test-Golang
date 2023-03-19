/*
Create time at 2023年3月18日0018下午 20:41:22
Create User at Administrator
*/

package initvariable

import (
	"JWT-Test-Golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func InitVariable() {
	var err error
	dsn := os.Getenv("Db")

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("conneted sql is failed")
	}
	//自动更新表结构
	Db.AutoMigrate(&models.User{})
}
