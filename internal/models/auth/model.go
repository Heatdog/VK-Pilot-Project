package auth

type ModelRequest struct {
	Login    string
	Password string
}

type ModelResponse struct {
	Token string `json:"token"`
}
