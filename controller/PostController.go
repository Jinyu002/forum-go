package controller

import (
	"forum1/common"
	"forum1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublishPost 发表帖子存入数据库
func PublishPost(c *gin.Context) {
	DB := common.GetDB()

	//获取参数
	params := model.PostParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400, "error": err.Error(),
		})
		return
	}

	//校验参数
	if len(params.Title) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": 421, "msg": "标题不能为空",
		})
		return
	}

	//查询用户id
	var user model.User
	DB.Where("username = ?", params.Username).First(&user)
	user_id := user.ID

	//创建帖子
	newPost := model.Post{
		User_id: user_id,
		Username: params.Username,
		Title:  params.Title,
		Content: params.Content,
	}
	DB.Create(&newPost)
}
