package createHolderUseCase

import (
	"fmt"
	holderEntity "holder-ms/core/entities/holder"
	"holder-ms/infra/database"
	repository "holder-ms/infra/database/repositories/mysqlRepositories"
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
