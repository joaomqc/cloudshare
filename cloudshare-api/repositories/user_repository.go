package repositories

import (
	"errors"

	"github.com/joaomqc/cloudshare/entities"
	"github.com/joaomqc/cloudshare/models"
)

var users []entities.User

// AddNewUser adds a new user
func AddNewUser(newUser models.User) {
	user := entities.User{
		Username: newUser.Username,
		Password: newUser.Password,
	}

	if users == nil {
		users = []entities.User{user}
	} else {
		users = append(users, user)
	}
}

// GetUsers returns all users
func GetUsers() []models.User {
	userInfos := make([]models.User, len(users))

	for i, u := range users {
		userInfos[i] = models.User{
			Username: u.Username,
		}
	}

	return userInfos
}

// GetUserByUsername returns user by username
func GetUserByUsername(username string) (models.User, error) {
	var userInfo models.User

	user, err := findUser(username)

	if err != nil {
		return userInfo, err
	}

	userInfo = models.User{
		Username: user.Username,
	}

	return userInfo, nil
}

// ValidateCredentials validated the user's credentials
func ValidateCredentials(username string, password string) (bool, error) {
	user, err := findUser(username)

	if err != nil {
		return false, err
	}

	return user.Password == password, nil
}

func findUser(username string) (entities.User, error) {
	var user entities.User

	if users == nil {
		return user, errors.New("no users")
	}

	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}

	return user, errors.New("user not found")
}
