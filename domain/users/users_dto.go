package users

import (
	"strings"

	"github.com/pt-abhishek/users-api/utils"
)

//User DTO
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate validates the data in DTO
func (user *User) Validate() *utils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	if strings.ToLower(user.Email) == "" {
		return utils.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
