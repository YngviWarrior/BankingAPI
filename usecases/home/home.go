package homeUseCase

import (
	repository "api-go/infra/database/repositories/mock"
)

type HomeUseCase struct {
	HomeRepository repository.MockRepositoryInterface
}
