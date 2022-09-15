package signupUsecase

import userEntity "go-api/core/user/entity"

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

	output.InternalStatus = 1

	return
}
