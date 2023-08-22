package models

type UserData struct {
	Email string `json:"email"`
}

type ResponseData struct {
	Data UserData `json:"data"`
}
