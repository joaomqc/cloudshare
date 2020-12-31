package views

// UserInputModel is the input user info
type UserInputModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserOutputModel is the output user info
type UserOutputModel struct {
	Username string `json:"username"`
}
