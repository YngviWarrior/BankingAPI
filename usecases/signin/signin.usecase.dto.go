package signinUsecase

type InputSignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}

type OutputSignInDto struct {
	Token string `json:"token"`
}
