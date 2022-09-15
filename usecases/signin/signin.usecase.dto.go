package signinUsecase

type InputSignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OutputSignInDto struct {
	Token string `json:"token"`
}
