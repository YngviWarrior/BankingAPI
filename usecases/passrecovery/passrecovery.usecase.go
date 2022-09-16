package passrecovery

import (
	userEntity "go-api/core/user"
	constants "go-api/infra"
	repository "go-api/infra/database/repositories/mysql"
	"time"
)

type PassRecoveryUsecase struct {
	UserRepository             repository.UserRepositoryInterface
	UserRecoveryPassRepository repository.UserRecoveryPassRepositoryInterface
	MailJobsRepository         repository.MailJobsRepositoryInterface
}

func (s *PassRecoveryUsecase) PassRecovery(input InputPassRecoveryDto) (output OutputPassRecoveryDto) {
	user := s.UserRepository.FindByEmail(input.Email)

	if (user != userEntity.User{}) {
		output.InternalStatus = 0
		return
	}

	lastInsert := s.UserRecoveryPassRepository.CreateRecoveryPass(user.Id, time.Now().Format("2006-01-02 15:04:05"), time.Now().Add(time.Hour*24).Format("2006-01-02 15:04:05"))

	if lastInsert == 0 {
		output.InternalStatus = 0
		return
	}

	if !s.MailJobsRepository.CreateJob(user.Id, constants.MODELO_USUARIO_RECUPERAR_SENHA, lastInsert, 0, 0, 0, 0, 0, "") {
		output.InternalStatus = 2
		return
	}

	output.InternalStatus = 1

	return
}
