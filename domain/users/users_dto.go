package users

import (
	"fmt"

	"github.com/pt-abhishek/users-api/utils"
)

var (
	usersDB = make(map[int64]*User)
)

//Get finds by id
func (user *User) Get() *utils.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return utils.NewResourceNotFoundError(fmt.Sprintf("user not found with ID: %d", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

//Save saves a user
func (user *User) Save() (*User, *utils.RestErr) {
	existingUser := usersDB[user.ID]
	if existingUser != nil {
		return nil, utils.NewBadRequestError(fmt.Sprintf("User with ID: %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return user, nil
}
