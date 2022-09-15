package signinUsecase

func (s *SignInUsecase) SignIn(input InputSignInDto) (output OutputSignInDto) {
	user := s.UserRepository.FindByEmail(input.Email)

	output.Token = user.Senha

	return
}
