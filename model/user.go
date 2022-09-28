package model

type UserCtxKeyType string

const UserCtxKey UserCtxKeyType = "user"

type User struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}
