package users

import (
	"strings"

	"github.com/pt-abhishek/users-api/utils"
)

//User DAO
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate validates the data in DAO
func (user *User) Validate() *utils.RestErr {
	if strings.TrimSpace(strings.ToLower(user.Email)) == "" {
		return utils.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
