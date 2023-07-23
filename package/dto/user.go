package dto

type User struct {
	FName    string `json:"fname"`
	LName    string `json:"lname"`
	Country  string `sjon:"country"`
	Pwd      string `json:"pwd"`
	UserName string `json:"userName"`
}
