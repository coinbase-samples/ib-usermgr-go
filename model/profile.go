package model

import "time"

type ProfileResponse struct {
	UserId      string    `json:"userId"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	LegalName   string    `json:"description"`
	UserName    string    `json:"username"`
	Roles       []string  `json:"roles"`
	Address     string    `json:"address"`
	DateOfBirth string    `json:"dateOfBirth"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UpdateProfileRequest struct {
	UserId      string    `json:"userId"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	LegalName   string    `json:"description"`
	UserName    string    `json:"username"`
	Roles       []string  `json:"roles"`
	Address     string    `json:"address"`
	DateOfBirth string    `json:"dateOfBirth"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
