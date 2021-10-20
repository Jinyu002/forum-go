package main

import (
	"forum1/common"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type User struct {
// 	gorm.Model
// 	Username string `gorm:"type:varchar(50); not null"`
// 	Password string `gorm:"type:char(32); not null"`
// 	Email    string `gorm:"type:varchar(50); not null"`
// 	// Birthday string `gorm:"type:date; not null"`
// 	Sex          int    `gorm:"type:tinyint(1); not null"`
// 	Address      string `gorm:"type:varchar(255); not null"`
// 	Post_number  int    `gorm:"type:int(4); not null"`
// 	Reply_number int    `gorm:"type:int(4); not null"`
// 	Status       int    `gorm:"type:tinyint(1); not null"`
// }

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
	// r.POST("/api/auth/register", func(c *gin.Context) {
	// 	//获取参数
	// 	username := c.PostForm("username")
	// 	password := c.PostForm("password")
	// 	confirm := c.PostForm("confirm")
	// 	email := c.PostForm("email")
	// 	//sex1 := c.PostForm("sex")
	// 	province := c.PostForm("province")
	// 	city := c.PostForm("city")
	// 	area := c.PostForm("area")
	// 	fmt.Println(province)
	// 	address := province + city + area
	// 	fmt.Println(address)
	// 	//数据验证
	// 	if len(password) < 6 {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"status": 422, "msg": "密码必须大于6位",
	// 		})
	// 		return
	// 	}

	// 	if confirm != password {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"status": 423, "msg": "两次输入的密码不一致",
	// 		})
	// 		return
	// 	}

	// 	if !VerifyUsername(username) {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"status": 424, "msg": "用户名格式不合法",
	// 		})
	// 		return
	// 	}

	// 	if !VerifyEmail(email) {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"status": 425, "msg": "邮箱格式不合法",
	// 		})
	// 		return
	// 	}

	// 	log.Println(username, password, email)

	// 	//判断用户名是否存在
	// 	if isUsernameExist(db, username) {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"status": 426, "msg": "用户名重复",
	// 		})
	// 		return
	// 	}

	// 	password = Md5(password)

	// 	//创建用户
	// 	newUser := User{
	// 		Username: username,
	// 		Password: password,
	// 		Email:    email,
	// 		Address:  address,
	// 	}
	// 	db.Create(&newUser)

	// 	//返回结果
	// 	c.JSON(200, gin.H{
	// 		"message": "注册成功",
	// 	})
	// })
	// r.POST("/api/auth/register", controller.Register)
	// listen and serve on 0.0.0.0:8080
}

// func VerifyEmail(email string) bool {
// 	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
// 	reg := regexp.MustCompile(pattern)
// 	return reg.MatchString(email)
// }

// func VerifyUsername(username string) bool {
// 	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,19}$`
// 	reg := regexp.MustCompile(pattern)
// 	return reg.MatchString(username)
// }

// func isUsernameExist(db *gorm.DB, username string) bool {
// 	var user User
// 	db.Where("username = ?", username).First(&user)
// 	if user.ID != 0 {
// 		return true
// 	}

// 	return false
// }

// func InitDB() *gorm.DB {
// 	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&model.User{})
// 	return db
// }

// func Md5(password string) string {
// 	m := md5.New()
// 	m.Write([]byte(password))
// 	res := hex.EncodeToString(m.Sum(nil))
// 	return res
// }
