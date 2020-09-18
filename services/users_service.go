package services

import (
	"github.com/pt-abhishek/users-api/domain/users"
	"github.com/pt-abhishek/users-api/utils"
)

//CreateUser creates user from db
func CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	var savedUser *users.User
	savedUser, err := user.Save()
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

//GetUser finds user by id
func GetUser(id int64) (*users.User, *utils.RestErr) {
	user := users.User{
		ID: id,
	}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser updates user
func UpdateUser(isPartial bool, user users.User) (*users.User, *utils.RestErr) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if err = currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

//DeleteUser deletes user with the given id
func DeleteUser(id int64) (*users.User, *utils.RestErr) {
	currentUser, err := GetUser(id)
	if err != nil {
		return nil, err
	}
	if err = currentUser.Delete(); err != nil {
		return nil, err
	}
	return currentUser, nil
}
