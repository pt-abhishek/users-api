package users

import "encoding/json"

//PublicUser limited fields
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//PrivateUser all except password
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

/*Marshall works when user and marshaled users have same JSON fields
but it should have more complexity as compared to individually copying fields
*/
func (user *User) Marshall(isPublic bool) interface{} {
	userJSON, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJSON, &publicUser)
		return publicUser
	}
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

//Marshall marhshals array of isers
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for idx, user := range users {
		result[idx] = user.Marshall(isPublic)
	}
	return result
}
