package utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	userNotFoundError = "no rows in result set"
)

//ParseError processes MySQL errors
func ParseError(err error) *RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), userNotFoundError) {
			return NewBadRequestError("no record matching the given id")
		}
		return NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		return NewBadRequestError("invalid data")
	}
	return NewInternalServerError("error processing request")
}
