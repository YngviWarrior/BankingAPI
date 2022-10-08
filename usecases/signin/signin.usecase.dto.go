package signinUsecase

type InputSignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

type OutputSignInDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
