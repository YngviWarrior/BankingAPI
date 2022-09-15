package signinUsecase

import (
	repository "go-api/infra/database/mysql"
)

type SignInUsecase struct {
	UserRepository *repository.UserRepository
}
