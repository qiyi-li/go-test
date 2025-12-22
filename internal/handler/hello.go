package handler

import (
	"errors"
	"fmt"
	"go-test/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInput struct {
	Name string `json:"name"`
}

func HelloHandler(c *gin.Context) {
	var userInput UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := store.User{
		Name: userInput.Name,
	}
	result := store.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Printf("✅ 新用户创建成功! 详细信息: %+v\n", user)
	c.JSON(201, user)
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		var users []store.User
		result := store.DB.Find(&users)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, users)
		return
	}
	var user store.User
	result := store.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, user)
}

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
	var user store.User
	result := store.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	user.Name = userInput.Name
	result = store.DB.Model(&user).Updates(store.User{Name: userInput.Name})
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, user)
}

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
	result := store.DB.Delete(&store.User{}, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "用户已删除"})
}
