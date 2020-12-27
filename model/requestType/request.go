package request

// User login structure
type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
