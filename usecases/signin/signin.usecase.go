package signinUsecase

import (
	repository "go-api/infra/database/repositories/mysql"
)

type SignInUsecase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignInUsecase) SignIn(input InputSignInDto) (output OutputSignInDto) {
	user := s.UserRepository.FindByEmail(input.Email)

	output.Token = user.Senha

	return
}
