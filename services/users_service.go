package services

import (
	"github.com/pt-abhishek/users-api/domain/users"
	"github.com/pt-abhishek/users-api/utils"
)

//CreateUser creates user from db
func CreateUser(user users.User) (*users.User, *utils.RestErr) {
	return &user, nil
}
