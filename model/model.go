/**
 * Copyright 2022 Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

type UserCtxKeyType string

const UserCtxKey UserCtxKeyType = "user"

type User struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}