package users

import (
	"strings"

	"github.com/pt-abhishek/users-api/utils"
)

//ActiveStatus is the status
const (
	ActiveStatus = "Active"
)

//User DTO
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users Slice of user
type Users []User

//Validate validates the data in DTO
func (user *User) Validate() *utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	if strings.ToLower(user.Email) == "" {
		return utils.NewBadRequestError("Invalid Email Address")
	}
	if user.Password == "" {
		return utils.NewBadRequestError("password needed")
	}
	return nil
}
