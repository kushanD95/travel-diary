package dto

type User struct {
	FName    string `json:"fname"  validate:"required,max=50"`
	LName    string `json:"lname"  validate:"required,max=50"`
	Country  string `json:"country"  validate:"required,max=50"`
	Pwd      string `json:"pwd"  validate:"required"`
	UserName string `json:"username"  validate:"required,max=25"`
}

type UserLogin struct {
	UserName string `json:"username"  validate:"required,max=25"`
	Pwd      string `json:"pwd"  validate:"required"`
}
