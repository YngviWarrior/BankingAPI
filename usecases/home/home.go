package homeUseCase

import (
	repository "api-user/infra/database/repositories/mock"
)

type HomeUseCase struct {
	HomeRepository repository.MockRepositoryInterface
}
