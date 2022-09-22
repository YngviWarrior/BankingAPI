package mysql

import (
	constants "api-user/core"
	"api-user/core/entities/balance"
	database "api-user/infra/database"
	"api-user/infra/utils"
	"fmt"
	"log"
	math "math/big"
	"time"
)

type UserBalanceRepository struct{}

type UserBalanceRepositoryInterface interface {
	InitialBalance(idUser uint64) bool
	ModifyBalance(idBalanceReceiver uint64, idBalanceOrigin uint64, value float64, idUser uint64, idRef int64, idRef2 int64) bool
}

func (UserBalanceRepository) ModifyBalance(idBalanceReceiver uint64, idBalanceOrigin uint64, value float64, idUser uint64, idRef int64, idRef2 int64) bool {
	conn := database.GetConnection()
	defer conn.Close()
	tx, _ := conn.Begin()

	var val float64
	err := tx.QueryRow(`
		SELECT valor
		FROM saldo_valor
		WHERE id_usuario = ? AND id = ?
		FOR UPDATE;
	`, idUser, idBalanceReceiver).Scan(&val)

	if err != nil {
		tx.Rollback()
		log.Println("UBRMB 01: ", err)
		return false
	}

	beforeValue := val

	newVal := math.NewFloat(value)
	bVal := math.NewFloat(beforeValue)
	bVal.Add(bVal, newVal)
	v, _ := bVal.Float64()

	afterValue := v

	res, err := tx.Exec(`
		UPDATE saldo_valor
		SET valor = ?, processing = 0
		WHERE id_usuario = ?
		AND id = ?
	`, afterValue, idUser, idBalanceReceiver)

	if err != nil {
		tx.Rollback()
		fmt.Println("UBRMB 2: " + err.Error())
	}

	affcRows, _ := res.RowsAffected()
	if affcRows == 0 {
		tx.Rollback()
		fmt.Println("UBRMB 3: ")
	}

	var b balance.Balance

	switch idBalanceOrigin {
	case constants.SALDO_ORIGEM_DEPOSIT:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.1: ")
			return false
		}

		b.IdPedido.Int64 = idRef

	case constants.SALDO_ORIGEM_BONUS_INDICATION:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.2: ")
			return false
		}

		b.IdBonusIndicacao.Int64 = idRef

	case constants.SALDO_ORIGEM_BONUS_INDICATION_TRADER:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.3: ")
			return false
		}

		b.IdBonusTrader.Int64 = idRef

	case constants.SALDO_ORIGEM_WITHDRAW, constants.SALDO_ORIGEM_WITHDRAW_CANCELED:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.4: ")
			return false
		}

		b.IdBtcPayment.Int64 = idRef

	case constants.SALDO_ORIGEM_GAME_BET_ADD_CREDIT_LOSE, constants.SALDO_ORIGEM_GAME_BET_REMOVE_CREDIT_WIN, constants.SALDO_ORIGEM_GAME_BET_REFUND, constants.SALDO_ORIGEM_GAME_BET_WIN, constants.SALDO_ORIGEM_GAME_BET:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.5: ")
			return false
		}

		b.IdBinaryOptionGameBet.Int64 = idRef

	case constants.SALDO_ORIGEM_FREE:
		b.IdDepositAdmin.Int64 = idRef

	case constants.SALDO_ORIGEM_DELETED_OPERATION:
		b.IdBinaryOptionGameBet.Int64 = idRef
		b.IdDepositAdmin.Int64 = idRef2

	case constants.SALDO_ORIGEM_CONVERSION:
		if idRef <= 0 {
			tx.Rollback()
			fmt.Println("UBRMB 4.6: ")
			return false
		}

		b.IdUserBalanceConvert.Int64 = idRef

	default:
		tx.Rollback()
		fmt.Println("UBRMB 4.7: ")
		return false

	}

	b.IdUsuario = idUser
	b.IdTipo = idBalanceReceiver
	b.IdOrigem = idBalanceOrigin
	b.Valor = utils.Trucate(value, 8)
	b.TotalAntes = uint64(utils.Trucate(beforeValue, 8))
	b.TotalDepois = uint64(utils.Trucate(afterValue, 8))
	b.DataRegistro = time.Now().Format("2006-01-02 15:04:05")

	var BI BalanceRepositoryInterface = &BalanceRepository{}
	if !BI.RegisterBalance(b, tx) {
		tx.Rollback()
		fmt.Println("UBRMB 4: ")
		return false
	}

	tx.Commit()

	return true
}

func (UserBalanceRepository) InitialBalance(idUser uint64) bool {
	conn := database.GetConnection()

	res, err := conn.Exec(`
	INSERT INTO saldo_valor (id, id_usuario, valor)
	SELECT id, ?, 0 FROM saldos_tipo
	`, idUser)

	if err != nil {
		log.Println("UBRIB 01: ", err)
		return false
	}

	defer conn.Close()

	lastId, _ := res.LastInsertId()

	if lastId == 0 {
		log.Println("UBRIB 02")
		// return false
	}

	return true
}
