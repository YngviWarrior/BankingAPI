package signupUsecase

type InputSignUpDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Sponsor  string `json:"sponsor"`
}

type OutputSignUpDto struct {
	InternalStatus int64
	Token          string `json:"token"`
}
