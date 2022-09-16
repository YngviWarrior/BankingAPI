package homeUseCase

import (
	repository "go-api/infra/database/repositories/mock"
)

type HomeUseCase struct {
	HomeRepository repository.MockRepositoryInterface
}
