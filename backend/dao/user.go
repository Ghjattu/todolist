package dao

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"size:255;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
}

type APIUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// BeforeCreate is a gorm hook that hashes user's password
// and escapes special characters in the user's username.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = string(hashedPassword)

	return nil
}

// CreateNewUser creates a new user.
func CreateNewUser(user *User) (*APIUser, error) {
	err := Db.Create(user).Error
	if err != nil {
		return &APIUser{}, err
	}

	returnedUser := &APIUser{
		ID:       user.ID,
		Username: user.Username,
	}

	return returnedUser, nil
}

// GetUerByUsername retrieves the specified user by the unique username.
// omitPassword: whether to return the user's password.
func GetUserByUsername(username string, omitPassword bool) (*APIUser, string, error) {
	user := &User{}
	if err := Db.Model(&User{}).Where("username = ?", username).First(user).Error; err != nil {
		return &APIUser{}, "", err
	}

	returnedUser := &APIUser{
		ID:       user.ID,
		Username: user.Username,
	}

	if omitPassword {
		return returnedUser, "", nil
	}
	return returnedUser, user.Password, nil
}
