package services

import (
	"github.com/pt-abhishek/users-api/domain/users"
	"github.com/pt-abhishek/users-api/utils"
)

//UserService is the single instance of the UserService interface
var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *utils.RestErr)
	GetUser(int64) (*users.User, *utils.RestErr)
	UpdateUser(bool, users.User) (*users.User, *utils.RestErr)
	Search(string) (users.Users, *utils.RestErr)
	DeleteUser(int64) (*users.User, *utils.RestErr)
}

//CreateUser creates user from db
func (u *userService) CreateUser(user users.User) (*users.User, *utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.ActiveStatus
	user.Password = utils.GetMD5(user.Password)

	var savedUser *users.User
	savedUser, err := user.Save()
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

//GetUser finds user by id
func (u *userService) GetUser(id int64) (*users.User, *utils.RestErr) {
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
func (u *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *utils.RestErr) {
	currentUser, err := UserService.GetUser(user.ID)
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
		if user.Status != "" {
			currentUser.Status = user.Status
		}

	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
		currentUser.Status = user.Status
	}

	if err = currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

//DeleteUser deletes user with the given id
func (u *userService) DeleteUser(id int64) (*users.User, *utils.RestErr) {
	currentUser, err := UserService.GetUser(id)
	if err != nil {
		return nil, err
	}
	if err = currentUser.Delete(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

//Search fetch by status
func (u *userService) Search(status string) (users.Users, *utils.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
