package models

type AuthRequest struct {
	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user_login"`
}

type AuthResponse struct {
	Success      bool   `json:"success"`
	UserApiToken string `json:"user_api_token"`
	UserID       int    `json:"user_id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Company      string `json:"company"`
	TimeZone     string `json:"time_zone"`
	Language     string `json:"language"`
}