package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Vikas208/social-media-api/helpers"
	"github.com/Vikas208/social-media-api/schemas"
	service "github.com/Vikas208/social-media-api/services"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var signupData schemas.SignupJsonSchema
	err := helpers.DecodeBodyToJson(w, r, &signupData)
	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Error(), mr.Status())
		} else {
			log.Printf("unknown error %s", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	response, err := service.Signup(w, signupData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	helpers.MarshalJson(response, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData schemas.LoginSchema
	err := helpers.DecodeBodyToJson(w, r, &loginData)
	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Error(), mr.Status())
		} else {
			log.Printf("unknown error %s", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	response, err := service.Login(loginData)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	helpers.MarshalJson(response, w)
}
