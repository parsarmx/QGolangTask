package responses

type RegisterResponse struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Status       string `json:"status"`
}

type CreatePostResponse struct {
	Created bool   `json:"created"`
	Status  string `json:"status"`
}

type UpdatePostResponse struct {
	Updated bool   `json:"updated"`
	Status  string `json:"status"`
}
