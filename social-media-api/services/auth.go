package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/Vikas208/social-media-api/helpers"
	"github.com/Vikas208/social-media-api/loaders"
	"github.com/Vikas208/social-media-api/models"
	"github.com/Vikas208/social-media-api/schemas"
	"gorm.io/gorm"
)

func Signup(w http.ResponseWriter, body schemas.SignupJsonSchema) (*helpers.Response, error) {
	var response *helpers.Response
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
				} else {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}()

	hahsPassword, err := helpers.HashPassword(body.Password)

	if err != nil {
		panic(err)
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: hahsPassword}

	if err := loaders.DB.Create(&user).Error; err != nil {

		panic(err)
	}

	response = &helpers.Response{
		Message: "Success",
		Code:    http.StatusOK,
	}
	return response, nil
}

func Login(body schemas.LoginSchema) (*helpers.Response, error) {
	var response *helpers.Response
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recover from: ", err)
			response = &helpers.Response{
				Message: "Error",
				Code:    http.StatusInternalServerError,
			}
		}
	}()

	// find email id
	var login models.User
	err := loaders.DB.Find(&login, "email = ?", body.Email)

	if err != nil {
		log.Panic(err)
	}

	// compare password

	isPasswordMatch := helpers.CheckPasswordHash(body.Password, login.Password)

	if !isPasswordMatch {
		response = &helpers.Response{
			Message: "Invalid Credentials",
			Code:    http.StatusUnauthorized,
		}
	}

	return response, nil
}
