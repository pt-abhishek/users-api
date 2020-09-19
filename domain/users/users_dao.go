package users

import (
	"fmt"

	"github.com/pt-abhishek/users-api/databases/mysql"
	"github.com/pt-abhishek/users-api/logger"
	"github.com/pt-abhishek/users-api/utils"
)

var (
	queryInsertUser   = "INSERT into users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUserByID  = "SELECT id, first_name, last_name, email, date_created, status from users WHERE id = ?;"
	queryUpdateUser   = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser   = "DELETE FROM users WHERE id = ?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status from users WHERE status = ?;"
)

//Get finds by id
func (user *User) Get() *utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUserByID)
	if err != nil {
		logger.Error("internal server error", err)
		return utils.NewInternalServerError("Some error occured")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return utils.ParseError(err)
	}
	return nil
}

//Save saves a user
func (user *User) Save() (*User, *utils.RestErr) {

	stmt, err := mysql.Client.Prepare(queryInsertUser)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = utils.GetNowDBFormat()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return nil, utils.ParseError(err)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, utils.ParseError(err)
	}

	user.ID = userID
	return user, nil
}

//Update Updates a user with a dto instance
func (user *User) Update() *utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryUpdateUser)
	if err != nil {
		return utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return utils.ParseError(err)
	}

	return nil
}

//Delete deletes user by id
func (user *User) Delete() *utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		return utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return utils.ParseError(err)
	}
	return nil

}

//FindByStatus finds a user by status string
func (user *User) FindByStatus(status string) ([]User, *utils.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var resUser User
		if err = rows.Scan(&resUser.ID, &resUser.FirstName, &resUser.LastName, &resUser.Email, &resUser.DateCreated, &resUser.Status); err != nil {
			return nil, utils.ParseError(err)
		}
		results = append(results, resUser)
	}
	if len(results) == 0 {
		return nil, utils.NewResourceNotFoundError(fmt.Sprintf("No Entities matching %s found", status))
	}
	return results, nil

}
