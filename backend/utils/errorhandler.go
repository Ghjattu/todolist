package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ErrorHandler(c *gin.Context, err error) {
	// default response
	code := http.StatusInternalServerError
	res := Response{
		Success: false,
		Code:    code,
		Message: "internal server error",
		Data:    nil,
	}

	// register or login input validation error
	if _, ok := err.(validator.ValidationErrors); ok {
		code = http.StatusBadRequest
		res.Code = code
		res.Message = "username or password is required"
	}

	// record not found error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		code = http.StatusNotFound
		res.Code = code
		res.Message = "record not found"
	}

	// password incorrect error
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		code = http.StatusUnauthorized
		res.Code = code
		res.Message = "password incorrect"
	}

	c.JSON(code, res)
}
