package users

//LoginRequest For external call to the api
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
