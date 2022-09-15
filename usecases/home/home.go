package homeUseCase

import (
	repository "go-api/infra/database/mock"
)

type HomeUseCase struct {
	HomeRepository *repository.MockRepository
}
