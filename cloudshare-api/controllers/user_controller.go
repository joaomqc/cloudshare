package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joaomqc/cloudshare/models"
	"github.com/joaomqc/cloudshare/repositories"
	"github.com/joaomqc/cloudshare/views"
)

// HandlePostUser creates new user
func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var user views.UserInputModel
	json.NewDecoder(r.Body).Decode(&user)

	newUser := models.User{
		Username: user.Username,
		Password: user.Password,
	}

	repositories.AddNewUser(newUser)
}

// HandleGetUsers gets all users
func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := repositories.GetUsers()

	outputUsers := make([]views.UserOutputModel, len(users))

	for i, u := range users {
		outputUsers[i] = views.UserOutputModel{
			Username: u.Username,
		}
	}

	json.NewEncoder(w).Encode(outputUsers)
}

// HandleGetUserByUsername gets user by username
func HandleGetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := repositories.GetUserByUsername(username)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	outputUser := views.UserOutputModel{
		Username: user.Username,
	}

	json.NewEncoder(w).Encode(outputUser)
}

// HandlePostLogin validates a user's credentials
func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	var user views.UserInputModel
	json.NewDecoder(r.Body).Decode(&user)

	valid, err := repositories.ValidateCredentials(user.Username, user.Password)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if !valid {
		http.Error(w, http.StatusText(401), 401)
	}
}
