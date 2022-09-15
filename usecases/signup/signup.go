package signupUsecase

import (
	repository "go-api/infra/database/mysql"
)

type SignUpUsecase struct {
	UserRepository *repository.UserRepository
}
