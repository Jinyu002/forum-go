package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"forum1/common"
	"forum1/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirm := c.PostForm("confirm")
	email := c.PostForm("email")
	//sex1 := c.PostForm("sex")
	province := c.PostForm("province")
	city := c.PostForm("city")
	area := c.PostForm("area")
	fmt.Println(province)
	address := province + city + area
	fmt.Println(address)
	//数据验证
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 422, "msg": "密码必须大于6位",
		})
		return
	}

	if confirm != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 423, "msg": "两次输入的密码不一致",
		})
		return
	}

	if !VerifyUsername(username) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 424, "msg": "用户名格式不合法",
		})
		return
	}

	if !VerifyEmail(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 425, "msg": "邮箱格式不合法",
		})
		return
	}

	log.Println(username, password, email)

	//判断用户名是否存在
	if isUsernameExist(DB, username) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 426, "msg": "用户名重复",
		})
		return
	}

	password = Md5(password)

	//创建用户
	newUser := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Address:  address,
	}
	DB.Create(&newUser)

	//返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func VerifyEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyUsername(username string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,19}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

func isUsernameExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func Md5(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
