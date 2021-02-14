package dto

// CreateUserRequest request dto to create new user
type CreateUserRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// OneUserResponse is the response of user information
type OneUserResponse struct {
	ID       string
	Username string `json:"username"`
}
