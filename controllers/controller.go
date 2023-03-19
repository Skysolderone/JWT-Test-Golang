/*
Create time at 2023年3月18日0018下午 20:36:21
Create User at Administrator
*/

package controllers

import (
	"JWT-Test-Golang/initvariable"
	"JWT-Test-Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	//获取用户结构
	var user models.User
	//绑定数据到创建的结构体
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}
	//给密码加密
	hashpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	user.Password = string(hashpass)
	//写进数据库
	initvariable.Db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"Result": "Success",
	})

}
func Login(c *gin.Context) {
	//获取User结构
	var user models.User
	//绑定json
	_ = c.ShouldBindJSON(&user)
	//在数据库里查找是否有该用户
	compass := user.Password
	err := initvariable.Db.Model(&user).Where("Email=?", user.Email).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "not found user",
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(compass))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "密码错误",
		})
	}
	//创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenstring, err := token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "token Signed error",
		})
	}
	//token写进cookie
	c.SetCookie("Authorization", tokenstring, 300, "", "", false, true)

	c.JSON(200, gin.H{
		"Success": "Success",
		"token":   tokenstring,
	})

}
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"user":    user,
		"message": "login in",
	})
}
