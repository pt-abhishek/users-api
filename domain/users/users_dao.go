package users

import (
	"github.com/pt-abhishek/users-api/databases/mysql"
	"github.com/pt-abhishek/users-api/utils"
)

var (
	queryInsertUser  = "INSERT into users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUserByID = "SELECT id, first_name, last_name, email, date_created from users WHERE id = ?;"
	queryUpdateUser  = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser  = "DELETE FROM users WHERE id = ?;"
)

//Get finds by id
func (user *User) Get() *utils.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetUserByID)
	if err != nil {
		return utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
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

	user.DateCreated = utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
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
