package signinUsecase

import (
	userEntity "api-user/core/entities/user"
	repository "api-user/infra/database/repositories/mysql"
	"api-user/infra/jwt"
	"api-user/infra/utils"
	"errors"
)

type SignInUsecase struct {
	UserRepository repository.UserRepositoryInterface
}

func (s *SignInUsecase) SignIn(input InputSignInDto) (output OutputSignInDto, err error) {
	user := s.UserRepository.FindByEmail(input.Email)
	encPass := utils.EncryptPassHS256(input.Password)

	if (user == userEntity.User{}) || encPass != user.Senha {
		err = errors.New("invalid email or password")
		return
	}

	var jwtInterface jwt.JwtInterface = &jwt.Jwt{}
	token, err := jwtInterface.GenerateJWT(input.IP)

	if err != nil {
		err = errors.New("internal error")
		return
	}

	output.Token = token

	return
}
