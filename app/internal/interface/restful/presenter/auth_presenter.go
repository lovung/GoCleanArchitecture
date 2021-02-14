package presenter

// RegisterRequest body of register request
type RegisterRequest struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}
