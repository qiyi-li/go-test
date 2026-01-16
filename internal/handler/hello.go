package handler

import (
	"errors"
	"fmt"
	"go-test/internal/service"
	"go-test/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ = store.User{}

type UserInput struct {
	Name string `json:"name"`
}

// HelloHandler 创建用户
// @Summary      创建用户
// @Description  创建新用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        input  body      UserInput  true  "用户信息"
// @Success      201    {object}  store.User
// @Failure      400    {object}  gin.H
// @Failure      500    {object}  gin.H
// @Router       /user [post]
func HelloHandler(c *gin.Context) {
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := service.CreateUser(userInput.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("✅ 新用户创建成功! 详细信息: %+v\n", user)
	c.JSON(201, user)
}

// GetUserHandler 获取用户详情
// @Summary      获取单个用户
// @Description  根据 ID 获取用户的详细信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  store.User  <-- 这里引用了 store
// @Failure      404  {object}  gin.H
// @Router       /user/{id} [get]
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// GetAllUserHandler 获取所有用户
// @Summary      获取所有用户
// @Description  获取所有用户的详细信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {array}  store.User
// @Failure      500  {object} gin.H
// @Router       /user [get]
func GetAllUserHandler(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// UpdateUserHandler 更新用户信息
// @Summary      更新用户信息
// @Description  根据 ID 更新用户的详细信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        user body      UserInput  true  "User Input"
// @Success      200  {object}  store.User
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /user/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	if c.Request.Method != "PUT" {
		c.JSON(405, gin.H{"error": "Method is not supported"})
		return
	}
	var userInput UserInput
	c.ShouldBindJSON(&userInput)
	user, err := service.UpdateUser(id, userInput.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Name = userInput.Name
	c.JSON(200, user)
}

// DeleteUserHandler 删除用户
// @Summary      删除用户
// @Description  根据 ID 删除用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  gin.H
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /user/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(405, gin.H{"error": "Method is not supported"})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	if err := service.DeleteUser(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "用户已删除"})
}
