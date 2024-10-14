package email

type AccessTokenRequest struct {
	ClientId     string `query:"clientId" validate:"required"`
	RefreshToken string `query:"refreshToken" validate:"required"`
}

type AuthResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`

	Error      string `json:"error"`
	ErrorDesc  string `json:"error_description"`
	ErrorCodes []int  `json:"error_codes"`
}
