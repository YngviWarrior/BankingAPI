package mysql

import (
	uniLevel "api-user/core/entities/uniLevel"
	database "api-user/infra/database"
	"log"
)

type UniLevelRepository struct{}

type UniLevelRepositoryInterface interface {
	SetPosition(uL uniLevel.UniLevel) bool
	FindByIdUsuario(idIndicator uint64) uniLevel.UniLevel
	UpdateByIdUsuario(idUsuario int64, idPai int64) bool
}

func (UniLevelRepository) FindByIdUsuario(idIndicator uint64) (u uniLevel.UniLevel) {

	conn := database.GetConnection()

	conn.QueryRow(`SELECT id_usuario, nivel FROM unilevel WHERE id_usuario = ?`, idIndicator).Scan(&u.IdUsuario, &u.Nivel)

	defer conn.Close()

	return
}

func (UniLevelRepository) UpdateByIdUsuario(idUsuario int64, idPai int64) bool {
	conn := database.GetConnection()

	res, err := conn.Exec(`
		UPDATE unilevel u 
		SET u.qtde_diretos = (u.qtde_diretos + 1) 
		WHERE u.id_usuario = ?
	`, idPai)

	if err != nil {
		log.Println("ULRUBIU 01: ", err)
		return false
	}

	affRows, _ := res.RowsAffected()

	if affRows == 0 {
		log.Println("ULRUBIU 02")
		return false
	}

	res, err = conn.Exec(`		
		UPDATE unilevel u 
		JOIN (
			SELECT linha_asc.id_ancestral
			FROM unilevel_rota linha_asc
			WHERE linha_asc.id_usuario = ?
		) as pai ON u.id_usuario = pai.id_ancestral
		SET u.qtde_rede = (u.qtde_rede + 1) 
	`, idUsuario)

	if err != nil {
		log.Println("ULRUBIU 03: ", err)
		return false
	}

	affRows, _ = res.RowsAffected()

	if affRows == 0 {
		log.Println("ULRUBIU 04")
		return false
	}

	defer conn.Close()

	return true
}

func (UniLevelRepository) SetPosition(uL uniLevel.UniLevel) bool {
	conn := database.GetConnection()

	var nivel int64
	conn.QueryRow(`SELECT nivel FROM unilevel WHERE id_usuario = ?`, uL.IdUsuario).Scan(&nivel)

	if nivel > 0 {
		return false
	}

	res, err := conn.Exec(`
		INSERT INTO unilevel(id_usuario, id_pai, nivel, qtde_diretos, qtde_rede) 
		VALUES(?, ?, ?, ?, ?)
	`, uL.IdUsuario.Int64, uL.IdPai.Int64, uL.Nivel.Int64, uL.QtdeDiretos, uL.QtdeRede)

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
