package mysql

import (
	database "api-user/infra/database"
	"log"
)

type UniLevelRouteRepository struct{}

type UniLevelRouteRepositoryInterface interface {
	CreateFromUniLevel(idUsuario int64, idPai int64) bool
	CreateUniLevelRoute(idUsuario int64, idPai int64, idNivel int64) bool
}

func (UniLevelRouteRepository) CreateUniLevelRoute(idUsuario int64, idPai int64, idNivel int64) bool {
	conn := database.GetConnection()

	res, err := conn.Exec(`
		INSERT INTO unilevel_rota (id_usuario,id_ancestral,nivel)
		VALUES(?, ?, ?)
	`, idUsuario, idPai, idNivel)

	if err != nil {
		log.Println("URC 01: ", err)
		return false
	}

	defer conn.Close()

	lastId, _ := res.LastInsertId()

	if lastId == 0 {
		log.Println("URC 02")
		return false
	}

	return true
}

func (UniLevelRouteRepository) CreateFromUniLevel(idUsuario int64, idPai int64) bool {
	conn := database.GetConnection()

	res, err := conn.Exec(`
		INSERT INTO unilevel_rota (id_usuario,id_ancestral,nivel)
		SELECT ?, linha_asc.id_ancestral, linha_asc.nivel
		FROM unilevel_rota linha_asc
		WHERE linha_asc.id_usuario = ?
	`, idUsuario, idPai)

	if err != nil {
		log.Println("URC 01: ", err)
		return false
	}

	defer conn.Close()

	lastId, _ := res.LastInsertId()

	if lastId == 0 {
		log.Println("URC 02")
		return false
	}

	return true
}
