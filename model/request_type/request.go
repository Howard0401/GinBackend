package request

// User login structure
type Login struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
