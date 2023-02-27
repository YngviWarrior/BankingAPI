package deleteHolderUseCase

import (
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
)

type DeleteHolderUsecase struct {
	Database         database.DatabaseInterface
	HolderRepository repository.HolderRepositoryInterface
}

func (u *DeleteHolderUsecase) DeleteHolder(input *InputDeleteHolderDto) (output OutputDeleteHolderDto, err error) {
	return
}
