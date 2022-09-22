package mysql

import (
	forbbiden "api-user/core/entities/forbbidenMailDomains"
	database "api-user/infra/database"
	"log"
)

type ForbbidenMailDomainsRepository struct{}

type ForbbidenMailDomainsRepositoryInterface interface {
	FindByDomain(domain string) bool
}

func (ForbbidenMailDomainsRepository) FindByDomain(domain string) bool {
	conn := database.GetConnection()

	var f forbbiden.ForbbidenMailDomains

	err := conn.QueryRow(`
		SELECT id, dominio
		FROM email_dominios_proibidos
		WHERE dominio = ?`, domain).Scan(&f.Id, &f.Domain)

	if err != nil {
		log.Println("FMDRFBD 01: ", err)
	}

	if f.Id > 0 {
		return false
	}

	defer conn.Close()

	return true
}
