package mysql

import (
	user "go-api/core/user"
	database "go-api/infra/database"

	"log"
)

type UserRepository struct{}

type UserRepositoryInterface interface {
	FindByEmail(email string) user.User
	Create() bool
	// FindAll() (list []string)
	// Update() bool
}

func (*UserRepository) Create() bool {
	return false
}

func (*UserRepository) FindByEmail(email string) (u user.User) {
	conn := database.GetConnection()

	err := conn.QueryRow(`
		SELECT id, usuario, senha, hash, id_indicador, data_cadastro, pin, email, admin, saldo, token, total_bonus, total_sacado, 
			total_deposito_btc, email_verificado, token_address, api_key, documento_verificado, genre, fullname, address, city, country_id, country_id_2,
			zip_code, phone, taxa_trade_percentual, birth_date, bot, id_qualification, approved_manager_status, trader_active, trader_percent_usage,
			id_trader, profile_image, withdraw_blocked, login_blocked_date_blocked, login_blocked_date_expire, login_blocked_date_reason, 
			bonus_indication_percent, total_deposit_balance_play, total_lose_balance_play, nome_documento, conta_excluida, id_idioma, identity, identity_user 
		FROM usuarios WHERE email = ?`, email).Scan(&u.Id, &u.Usuario, &u.Senha, &u.Hash, &u.IdIndicador, &u.DataCadastro,
		&u.Pin, &u.Email, &u.Admin, &u.Saldo, &u.Token, &u.TotalBonus, &u.TotalSacado, &u.TotalDepositoBtc, &u.EmailVerificado,
		&u.TokenAddress, &u.ApiKey, &u.DocumentoVerificado, &u.Genre, &u.Fullname, &u.Address, &u.City, &u.CountryId, &u.CountryId2,
		&u.ZipCode, &u.Phone, &u.TaxaTradePercentual, &u.BirthDate, &u.Bot, &u.IdQualification, &u.ApprovedManagerStatus, &u.TraderActive,
		&u.TraderPercentUsage, &u.IdTrader, &u.ProfileImage, &u.WithdrawBlocked, &u.LoginBlockedDateBlocked, &u.LoginBlockedDateExpire,
		&u.LoginBlockedDateReason, &u.BonusIndicationPercent, &u.TotalDepositBalancePlay, &u.TotalLoseBalancePlay, &u.NomeDocumento,
		&u.ContaExcluida, &u.IdIdioma, &u.Identity, &u.IdentityUser)

	if err != nil {
		log.Println("UFBE 01: ", err)
	}

	defer conn.Close()

	return
}
