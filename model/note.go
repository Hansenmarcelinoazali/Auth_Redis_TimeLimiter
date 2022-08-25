package model

// "github.com/jinzhu/gorm"

type Users struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	IsLogin      bool   `json:"is_login"`
	RefreshToken string `json:"refresh_token"`
}



