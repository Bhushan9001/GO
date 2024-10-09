package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Bhushan9001/GO_CRUD/config"
	"github.com/Bhushan9001/GO_CRUD/internal/models"
	"github.com/Bhushan9001/GO_CRUD/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid Payload")
		return
	}

	if user.Password == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "Password is Required")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user.Password = string(hashedPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)
}

func Signin(w http.ResponseWriter, r *http.Request) {

	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid Payload")
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "Email and Password are Required")
		return
	}

	var user models.User

	result := config.DB.Where("email = ?", loginRequest.Email).First(&user)

	if result.Error != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	response := struct {
		Token string      `json:"token"`
		User  models.User `json:"user"`
	}{
		Token: token,
		User:  user,
	}

	utils.RespondJSON(w, http.StatusOK, response)

}
