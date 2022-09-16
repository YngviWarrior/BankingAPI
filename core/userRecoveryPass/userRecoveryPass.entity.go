package userrecoverypass

import "database/sql"

type UserRecoveryPass struct {
	Id            uint64         `json:"id"`
	IdUsuario     uint64         `json:"id_usuario"`
	Code          string         `json:"code"`
	NovaSenhaHash sql.NullString `json:"nova_senha_hash"`
	DataAlterado  sql.NullString `json:"data_alterado"`
	DataRegistro  string         `json:"data_registro"`
	DataExpira    string         `json:"data_expira"`
}
