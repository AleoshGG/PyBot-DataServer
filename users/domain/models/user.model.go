package models

type User struct {
	User_id    int    `json:"user_id"`
	First_name string `json:"fist_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}