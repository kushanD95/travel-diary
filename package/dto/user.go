package dto

type CreateUserResponse struct {
	Status   int
	Message  string
	Username string
}

type UserLogin struct {
	UserName string
	Pwd      string
}

type LoginUserResponse struct {
	Status   int
	Message  string
	Username string
	Token    string
}
