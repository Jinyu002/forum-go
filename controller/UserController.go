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
	// useame := c.PostForm("username")
	// password := c.PostForm("password"
	// confirm := c.PostForm("confirm")
	// email := c.PostForm("email")
	// //sex1 := c.PostForm("sex")
	// province := c.PostForm("pronce")
	// city := c.PostForm("city")
	// area := c.PostForm("area")
	// fmt.Println(province)
	// address := province +ity + area
	// fmt.Println(address)
	params := model.UserParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "error": err.Error(),
		})
		return
	}
	//数据验证
	if len(params.Password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 422, "msg": "密码必须大于6位",
		})
		return
	}

	if params.Confirm != params.Password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 423, "msg": "两次输入的密码不一致",
		})
		return
	}

	if !VerifyUsername(params.Username) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 424, "msg": "用户名格式不合法",
		})
		return
	}

	if !VerifyEmail(params.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 425, "msg": "邮箱格式不合法",
		})
		return
	}

	log.Println(params.Username, params.Password, params.Email)

	//判断用户名是否存在
	if isUsernameExist(DB, params.Username) {
		fmt.Println(isUsernameExist(DB, params.Username))
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 426, "msg": "用户名重复",
		})
		return
	}

	password := Md5(params.Password)
	address := params.Province + params.City + params.Area

	//创建用户
	newUser := model.User{
		Username: params.Username,
		Password: password,
		Email:    params.Email,
		Birthday: params.Birthday,
		Sex:      params.Sex,
		Address:  address,
	}
	DB.Create(&newUser)

	//返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func  Login(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.UserParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "error": err.Error(),
		})
		return
	}

	//判断用户名是否存在
	var user model.User
	DB.Where("username = ?", params.Username).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 422, "msg": "用户名不存在",
		})
		return
	}

	//判断密码是否正确
	password := Md5(params.Password)
	if user.Password != password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 423, "msg": "密码错误",
		})
	}

	//发放token
	token := params.Username

	//返回结果
	c.JSON(200, gin.H{
		"status": 1, "msg": "登录成功", "data": gin.H{"token": token},
	})
}

// Ban 封禁用户
func Ban(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.UserParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": 400, "error": err.Error(),
		})
		return
	}

	//封禁操作
	var user model.User
	DB.Model(&user).Find(&user, "username = ?", params.Username).UpdateColumn("status", 1)

	//返回结果
	c.JSON(200, gin.H{
		"status": 1, "msg": "封禁成功",
	})

}

// Unban 解封用户
func Unban(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.UserParams{}
	if err := c.ShouldBind(&params); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "err": err.Error(),
		})
		return
	}

	//解封操作
	var user model.User
	DB.Model(&user).Find(&user, "username = ?", params.Username).UpdateColumn("status", 0)

	//返回结果
	c.JSON(200, gin.H{
		"status": 1, "msg": "解封成功",
	})
}

// Number 用户发帖回帖数量
func Number(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.UserParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "err": err.Error(),
		})
		return
	}

	//查询操作
	var user model.User
	DB.Find(&user, "username = ?", params.Username)

	//返回数据
	c.JSON(200, gin.H{
		"status": 1,  "msg": "查询成功", "data":gin.H{"post_number": user.Post_number, "reply_number": user.Reply_number},
	})
}

// ListUsers 管理员处展示所有用户
func ListUsers(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.QueryParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "error": err.Error(),
		})
		return
	}
	limit := params.Limit
	offset := (params.Page - 1) * params.Limit

	//查询操作
	var user []model.User
	DB.Limit(limit).Offset(offset).Find(&user)

	//返回数据
	c.JSON(200, gin.H{
		"status": 1, "msg": "查询成功", "data":user,
	})
}

func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
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
