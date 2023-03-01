package verifyHolderUseCase

import (
	holderEntity "api-go/core/entities/holder"
	"api-go/infra/database"
	repository "api-go/infra/database/repositories/mysqlRepositories"
	"fmt"
	"log"
)

type VerifyHolderUsecase struct {
	Database         database.DatabaseInterface
	HolderRepository repository.HolderRepositoryInterface
}

func (c *VerifyHolderUsecase) VerifyHolder(input *InputVerifyHolderDto) (output OutputVerifyHolderDto, err error) {
	conn := c.Database.CreateConnection()

	h := c.HolderRepository.FindByColumn(nil, conn, "cpf", input.CPF)

	if (h == holderEntity.Holder{}) {
		conn.Close()
		err = fmt.Errorf("holder not found")
		return
	}

	if h.Verified {
		conn.Close()
		err = fmt.Errorf("holder has already been verified")
		return
	}

	if !c.HolderRepository.Verify(nil, conn, h.Holder) {
		log.Println("CHUC 01")
		conn.Close()
		return
	}

	conn.Close()
	return
}
