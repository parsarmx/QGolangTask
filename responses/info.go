package responses

type RegisterResponse struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Status       string `json:"status"`
}
