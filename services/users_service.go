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
