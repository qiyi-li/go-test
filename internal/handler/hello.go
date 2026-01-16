package handler

import (
	"errors"
	"fmt"
	"go-test/internal/service"
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
	user, err := service.CreateUser(userInput.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("✅ 新用户创建成功! 详细信息: %+v\n", user)
	c.JSON(201, user)
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		users, err := service.GetUser("")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
		return
	}
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
