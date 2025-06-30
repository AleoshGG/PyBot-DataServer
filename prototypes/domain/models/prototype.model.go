package models

type Prototype struct {
	Prototype_id   int    `json:"prototype_id"`
	Prototype_name string `json:"prototype_name"`
	Model          string `json:"model"`
	User_id        int    `json:"user_id"`
}