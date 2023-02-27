package createHolderUseCase

import (
	holderEntity "api-go/core/entities/holder"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"fmt"
	"log"
)

type CreateHolderUsecase struct {
	Database         database.DatabaseInterface
	HolderRepository repository.HolderRepositoryInterface
}

func (c *CreateHolderUsecase) CreateHolder(input *InputCreateHolderDto) (output OutputCreateHolderDto, err error) {
	conn := c.Database.CreateConnection()

	h := c.HolderRepository.FindByColumn(nil, conn, "cpf", input.CPF)

	if (h != holderEntity.Holder{}) {
		err = fmt.Errorf("holder already exists")
		conn.Close()
		return
	}

	h.FullName = input.FullName
	h.CPF = input.CPF

	if !c.HolderRepository.Create(nil, conn, h) {
		log.Println("CHUC 01")
		conn.Close()
		return
	}

	conn.Close()
	return
}
