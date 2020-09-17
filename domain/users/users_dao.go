package users

import (
	"fmt"
	"strings"

	"github.com/pt-abhishek/users-api/databases/mysql"
	"github.com/pt-abhishek/users-api/utils"
)

var (
	usersDB          = make(map[int64]*User)
	queryInsertUser  = "INSERT into users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	indexUniqueEmail = "email_UNIQUE"
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

	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return nil, utils.NewBadRequestError(fmt.Sprintf("Email id %s already exists", user.Email))
		}
		return nil, utils.NewInternalServerError(fmt.Sprintf("Error inserting into the database %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, utils.NewInternalServerError("Error fetching the last insert id")
	}

	user.ID = userID
	return user, nil
}
