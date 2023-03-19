/*
Create time at 2023年3月18日0018下午 20:35:41
Create User at Administrator
*/

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
