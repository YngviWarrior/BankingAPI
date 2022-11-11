package signinUsecase

import (
	userEntity "api-go/core/entities/user"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysql"
	"api-go/infra/jwt"
	"api-go/infra/utils"
	"errors"
)

type SignInUsecase struct {
	Database       database.DatabaseInterface
	UserRepository repository.UserRepositoryInterface
}

func (s *SignInUsecase) SignIn(input *InputSignInDto) (output OutputSignInDto, err error) {
	tx, conn := s.Database.CreateConnection()

	user := s.UserRepository.FindByColumn(tx, "email", input.Email)
	encPass := utils.EncryptPassHS256(input.Password)

	if (user == userEntity.User{}) || encPass != user.Senha {
		tx.Rollback()
		err = errors.New("invalid email or password")
		return
	}

	var jwtInterface jwt.JwtInterface = &jwt.Jwt{}
	accessToken, refreshToken, err := jwtInterface.GenerateJWT(user.Id, user.Admin, input.IP)

	if err != nil {
		tx.Rollback()
		err = errors.New("internal error")
		return
	}

	output.AccessToken = accessToken
	output.RefreshToken = refreshToken

	tx.Commit()
	conn.Close()
	return
}
