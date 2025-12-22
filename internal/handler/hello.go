package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-test/internal/store"
	"net/http"

	"gorm.io/gorm"
)

type UserInput struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var userInput UserInput
	json.NewDecoder(r.Body).Decode(&userInput)
	if userInput.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	user := store.User{
		Name: userInput.Name,
	}
	result := store.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "保存失败: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("✅ 新用户创建成功! 详细信息: %+v\n", user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created 是更标准的“创建成功”状态码
	json.NewEncoder(w).Encode(user)
	// ch := make(chan int, 1)
	// go func() {
	// 	for i := range 3 {
	// 		time.Sleep(5 * time.Second)
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// select {
	// case msg := <-ch:
	// 	fmt.Println(msg)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("timeout")
	// 	http.Error(w, "超时啦", http.StatusGatewayTimeout)
	// 	return
	// }
	// for msg := range ch {
	// 	fmt.Println(msg)
	// }
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		var users []store.User
		result := store.DB.Find(&users)
		if result.Error != nil {
			http.Error(w, "获取用户列表失败: "+result.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	var user store.User
	result := store.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "用户不存在", http.StatusNotFound)
			return
		}
		http.Error(w, "获取用户失败: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	var userInput UserInput
	json.NewDecoder(r.Body).Decode(&userInput)
	var user store.User
	result := store.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "用户不存在", http.StatusNotFound)
			return
		}
		http.Error(w, "获取用户失败: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	user.Name = userInput.Name
	result = store.DB.Model(&user).Updates(store.User{Name: userInput.Name})
	if result.Error != nil {
		http.Error(w, "更新用户失败: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	result := store.DB.Delete(&store.User{}, id)
	if result.Error != nil {
		http.Error(w, "删除用户失败: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
