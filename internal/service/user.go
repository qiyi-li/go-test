package service

import (
	"errors"
	"go-test/internal/store"

	"gorm.io/gorm"
)

func CreateUser(name string) (*store.User, error) {
	user := store.User{Name: name}
	result := store.DB.Create(&user)
	return &user, result.Error
}

func GetUser(id string) (*store.User, error) {
	var user store.User
	if id == ""{
		var users []store.User
		result := store.DB.Find(&users)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	result := store.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UpdateUser(id string, newName string) (*store.User, error) {
	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}
	if err := store.DB.Model(user).Update("Name", newName).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id string) error {
	return store.DB.Delete(&store.User{}, id).Error
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
