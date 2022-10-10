package mysql

import (
	user "api-go/core/entities/user"
	"api-go/infra/utils"
	"database/sql"

	"log"
)

type UserRepository struct{}

type UserRepositoryInterface interface {
	FindByColumn(tx *sql.Tx, column string, value any) (u user.User)
	Create(tx *sql.Tx, u user.User) user.User
	UpdateDynamically(tx *sql.Tx, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool
}

func (*UserRepository) UpdateDynamically(tx *sql.Tx, updateFields []string, updatefieldValues []any, wherecolumns []string, wherevalues []any, paginationValues []any, order string) bool {
	_, wheres, updates := utils.QueryFormatter(updateFields, updatefieldValues, wherecolumns, wherevalues, paginationValues, order)
	query := `UPDATE usuarios SET ` + updates + wheres

	_, err := tx.Exec(query)

	if err != nil {
		log.Println("URUBC 01: ", err)
		return false
	}

	return true
}

func (*UserRepository) UpdateSenha(tx *sql.Tx, password string, userId uint64) bool {
	res, err := tx.Exec(`
		UPDATE usuarios SET senha = ? WHERE id = ?
	`, password, userId)

	if err != nil {
		log.Println("URUS 01: ", err)
		return false
	}

	affcRows, _ := res.RowsAffected()

	if affcRows == 0 {
		log.Println("URUS 02")
		return false
	}

	return true
}

func (*UserRepository) UpdateEmailVerify(tx *sql.Tx, userId uint64) bool {

	res, err := tx.Exec(`
		UPDATE usuarios SET email_verificado = 1 WHERE id = ?
	`, userId)

	if err != nil {
		log.Println("URC 01: ", err)
		return false
	}

	affcRows, _ := res.RowsAffected()

	if affcRows == 0 {
		log.Println("URC 02")
		return false
	}

	return true
}

func (*UserRepository) Create(tx *sql.Tx, user user.User) (u user.User) {

	res, err := tx.Exec(`
		INSERT INTO usuarios(hash, email, senha, id_indicador, bonus_indication_percent, taxa_trade_percentual, data_cadastro) 
		VALUES(?, ?, ?, ?, ?, ?, ?)
	`, user.Hash, user.Email.String, user.Senha, user.IdIndicador, user.BonusIndicationPercent.Float64, user.TaxaTradePercentual, user.DataCadastro.String)

	if err != nil {
		log.Println("URC 01: ", err)
		return
	}

	lastId, _ := res.LastInsertId()

	if lastId == 0 {
		log.Println("URC 02")
		return
	}

	err = tx.QueryRow(`
		SELECT id, usuario, senha, hash, id_indicador, data_cadastro, pin, email, admin, saldo, token, total_bonus, total_sacado, 
			total_deposito_btc, email_verificado, token_address, api_key, documento_verificado, genre, fullname, address, city, country_id, country_id_2,
			zip_code, phone, taxa_trade_percentual, birth_date, bot, id_qualification, approved_manager_status, trader_active, trader_percent_usage,
			id_trader, profile_image, withdraw_blocked, login_blocked_date_blocked, login_blocked_date_expire, login_blocked_date_reason, 
			bonus_indication_percent, total_deposit_balance_play, total_lose_balance_play, nome_documento, conta_excluida, id_idioma, identity, identity_user 
		FROM usuarios WHERE id = ?`, lastId).Scan(&u.Id, &u.Usuario, &u.Senha, &u.Hash, &u.IdIndicador, &u.DataCadastro,
		&u.Pin, &u.Email, &u.Admin, &u.Saldo, &u.Token, &u.TotalBonus, &u.TotalSacado, &u.TotalDepositoBtc, &u.EmailVerificado,
		&u.TokenAddress, &u.ApiKey, &u.DocumentoVerificado, &u.Genre, &u.Fullname, &u.Address, &u.City, &u.CountryId, &u.CountryId2,
		&u.ZipCode, &u.Phone, &u.TaxaTradePercentual, &u.BirthDate, &u.Bot, &u.IdQualification, &u.ApprovedManagerStatus, &u.TraderActive,
		&u.TraderPercentUsage, &u.IdTrader, &u.ProfileImage, &u.WithdrawBlocked, &u.LoginBlockedDateBlocked, &u.LoginBlockedDateExpire,
		&u.LoginBlockedDateReason, &u.BonusIndicationPercent, &u.TotalDepositBalancePlay, &u.TotalLoseBalancePlay, &u.NomeDocumento,
		&u.ContaExcluida, &u.IdIdioma, &u.Identity, &u.IdentityUser)

	if err != nil {
		log.Println("URC 03: ", err)
	}

	return
}

func (*UserRepository) FindByColumn(tx *sql.Tx, colunm string, value any) (u user.User) {

	err := tx.QueryRow(`
		SELECT id, usuario, senha, hash, id_indicador, data_cadastro, pin, email, admin, saldo, token, total_bonus, total_sacado, 
			total_deposito_btc, email_verificado, token_address, api_key, documento_verificado, genre, fullname, address, city, country_id, country_id_2,
			zip_code, phone, taxa_trade_percentual, birth_date, bot, id_qualification, approved_manager_status, trader_active, trader_percent_usage,
			id_trader, profile_image, withdraw_blocked, login_blocked_date_blocked, login_blocked_date_expire, login_blocked_date_reason, 
			bonus_indication_percent, total_deposit_balance_play, total_lose_balance_play, nome_documento, conta_excluida, id_idioma, identity, identity_user 
		FROM usuarios WHERE `+colunm+` = ?`, value).Scan(&u.Id, &u.Usuario, &u.Senha, &u.Hash, &u.IdIndicador, &u.DataCadastro,
		&u.Pin, &u.Email, &u.Admin, &u.Saldo, &u.Token, &u.TotalBonus, &u.TotalSacado, &u.TotalDepositoBtc, &u.EmailVerificado,
		&u.TokenAddress, &u.ApiKey, &u.DocumentoVerificado, &u.Genre, &u.Fullname, &u.Address, &u.City, &u.CountryId, &u.CountryId2,
		&u.ZipCode, &u.Phone, &u.TaxaTradePercentual, &u.BirthDate, &u.Bot, &u.IdQualification, &u.ApprovedManagerStatus, &u.TraderActive,
		&u.TraderPercentUsage, &u.IdTrader, &u.ProfileImage, &u.WithdrawBlocked, &u.LoginBlockedDateBlocked, &u.LoginBlockedDateExpire,
		&u.LoginBlockedDateReason, &u.BonusIndicationPercent, &u.TotalDepositBalancePlay, &u.TotalLoseBalancePlay, &u.NomeDocumento,
		&u.ContaExcluida, &u.IdIdioma, &u.Identity, &u.IdentityUser)

	if err != nil {
		log.Println("URFBI 01: ", err)
	}

	return
}
