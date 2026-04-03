package models

type UserResponse struct {
	UserID    int  `json:"user_id"`
	IsCorrect bool `json:"is_correct"`
}