package userEntity

import "database/sql"

type User struct {
	Id                      uint64          `json:"id"`
	Usuario                 sql.NullString  `json:"usuario"`
	Senha                   string          `json:"senha"`
	Hash                    string          `json:"hash"`
	IdIndicador             uint64          `json:"id_indicador"`
	DataCadastro            sql.NullString  `json:"data_cadastro"`
	Pin                     sql.NullString  `json:"pin"`
	Email                   sql.NullString  `json:"email"`
	Admin                   bool            `json:"admin"`
	Saldo                   sql.NullFloat64 `json:"saldo"`
	Token                   sql.NullFloat64 `json:"token"`
	TotalBonus              sql.NullFloat64 `json:"total_bonus"`
	TotalSacado             sql.NullFloat64 `json:"total_sacado"`
	TotalDepositoBtc        sql.NullFloat64 `json:"total_deposito_btc"`
	EmailVerificado         bool            `json:"email_verificado"`
	TokenAddress            sql.NullString  `json:"token_address"`
	ApiKey                  sql.NullString  `json:"api_key"`
	DocumentoVerificado     bool            `json:"documento_verificado"`
	Genre                   sql.NullString  `json:"genre"`
	Fullname                sql.NullString  `json:"fullname"`
	Address                 sql.NullString  `json:"address"`
	City                    sql.NullString  `json:"city"`
	CountryId               sql.NullInt64   `json:"country_id"`
	CountryId2              sql.NullInt64   `json:"country_id_2"`
	ZipCode                 sql.NullString  `json:"zip_code"`
	Phone                   sql.NullString  `json:"phone"`
	TaxaTradePercentual     float64         `json:"taxa_trade_percentual"`
	BirthDate               sql.NullString  `json:"birth_date"`
	Bot                     bool            `json:"bot"`
	IdQualification         bool            `json:"id_qualification"`
	ApprovedManagerStatus   bool            `json:"approved_manager_status"`
	TraderActive            bool            `json:"trader_active"`
	TraderPercentUsage      float64         `json:"trader_percent_usage"`
	IdTrader                sql.NullInt64   `json:"id_trader"`
	ProfileImage            sql.NullString  `json:"profile_image"`
	WithdrawBlocked         sql.NullBool    `json:"withdraw_blocked"`
	LoginBlockedDateBlocked sql.NullString  `json:"login_blocked_date_blocked"`
	LoginBlockedDateExpire  sql.NullString  `json:"login_blocked_date_expire"`
	LoginBlockedDateReason  sql.NullString  `json:"login_blocked_date_reason"`
	BonusIndicationPercent  sql.NullFloat64 `json:"bonus_indication_percent"`
	TotalDepositBalancePlay float64         `json:"total_deposit_balance_play"`
	TotalLoseBalancePlay    float64         `json:"total_lose_balance_play"`
	NomeDocumento           sql.NullString  `json:"nome_documento"`
	ContaExcluida           sql.NullBool    `json:"conta_excluida"`
	IdIdioma                sql.NullBool    `json:"id_idioma"`
	Identity                sql.NullString  `json:"identity"`
	IdentityUser            sql.NullString  `json:"identity_user"`
}
