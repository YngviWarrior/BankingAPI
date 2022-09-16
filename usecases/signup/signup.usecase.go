package signupUsecase

import (
	userEntity "go-api/core/user"
	repository "go-api/infra/database/repositories/mysql"
)

type SignUpUsecase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignUpUsecase) SignUp(input InputSignUpDto) (output OutputSignUpDto) {
	user := s.UserRepository.FindByEmail(input.Email)

	if (user != userEntity.User{}) {
		output.InternalStatus = 0
		return
	}

	if !s.UserRepository.Create() {
		output.InternalStatus = 2
		return
	}

	//Ger Pass

	output.InternalStatus = 1

	return
}
