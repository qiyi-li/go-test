package store
import "gorm.io/gorm"

type User struct {
	Name string `json:"name"`
	gorm.Model
}