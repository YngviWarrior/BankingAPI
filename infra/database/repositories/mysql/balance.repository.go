package mysql

import (
	balance "api-user/core/entities/balance"
	database "api-user/infra/database"
	"database/sql"
	"log"
)

type BalanceRepository struct{}

type BalanceRepositoryInterface interface {
	RegisterBalance(balance.Balance, *sql.Tx) bool
}

func (BalanceRepository) RegisterBalance(b balance.Balance, tx *sql.Tx) bool {
	var conn *sql.DB
	if tx == nil {
		conn = database.GetConnection()
		defer conn.Close()
		tx, _ = conn.Begin()
	}

	res, err := tx.Exec(`
		INSERT INTO saldos (id_usuario, id_tipo, id_origem, valor, total_antes, total_depois, data_registro, 
			id_pedido, id_bonus_indicacao, id_bonus_trader, id_btc_payment, id_binary_option_game_bet, 
			id_deposit_admin, id_user_balance_convert)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, b.IdUsuario, b.IdTipo, b.IdOrigem, b.Valor, b.TotalAntes, b.TotalDepois, b.DataRegistro, b.IdPedido,
		b.IdBonusIndicacao, b.IdBonusTrader, b.IdBtcPayment, b.IdBinaryOptionGameBet,
		b.IdDepositAdmin, b.IdUserBalanceConvert)

	if err != nil {
		tx.Rollback()
		log.Println("URC 01: ", err)
		return false
	}

	lastId, _ := res.LastInsertId()

	if lastId == 0 {
		tx.Rollback()
		log.Println("URC 02")
		return false
	}

	return true
}
