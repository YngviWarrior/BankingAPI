package mysql

import (
	userGov "api-user/core/entities/userGovernmentIds"
	database "api-user/infra/database"
	"log"
)

type UserGovernmentIdsRepository struct{}

type UserGovernmentIdsRepositoryInterface interface {
	FindByIdUsuarioAndName(idUser uint64, name []string) (list []*userGov.UserGovernmentIds)
}

func (UserGovernmentIdsRepository) FindByIdUsuarioAndName(idUser uint64, name []string) (list []*userGov.UserGovernmentIds) {
	conn := database.GetConnection()

	if len(name) == 1 {
		var u userGov.UserGovernmentIds
		err := conn.QueryRow(`
			SELECT value
			FROM usuarios_government_ids
			WHERE id_usuario = ? AND name = ?
		`, idUser, name[0]).Scan(&u.Value)

		if err != nil {
			log.Println("UBRFBIYAB 01: ", err)
			return
		}

		defer conn.Close()

		list = append(list, &u)
	} else {
		var names string

		res, err := conn.Query(`
			SELECT name, value
			FROM usuarios_government_ids
			WHERE id_usuario = ? AND name IN (`+names+`)
		`, idUser)

		if err != nil {
			log.Println("UBRFBIYAB 02: ", err)
			return
		}

		defer conn.Close()

		for res.Next() {
			var u userGov.UserGovernmentIds
			err := res.Scan(&u.Name, &u.Value)

			if err != nil {
				return
			}

			list = append(list, &u)
		}

	}

	return
}
