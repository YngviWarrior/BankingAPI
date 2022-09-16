package mailjobs

import "database/sql"

// table model tarefas
type MailJobs struct {
	Id                   uint64         `json:"id"`
	IdTipo               bool           `json:"id_tipo"`
	IdStatus             bool           `json:"id_status"`
	IdUsuario            uint64         `json:"id_usuario"`
	DataRegistro         string         `json:"data_registro"`
	DataUltimaTentativa  sql.NullString `json:"data_ultima_tentativa"`
	DataProximaTentativa string         `json:"data_proxima_tentativa"`
	DataCompleto         sql.NullString `json:"data_completo"`
	Ref1                 sql.NullString `json:"ref1"`
	Ref2                 sql.NullString `json:"ref2"`
	Ref3                 sql.NullString `json:"ref3"`
	Ref4                 sql.NullString `json:"ref4"`
	Ref5                 sql.NullString `json:"ref5"`
	Ref6                 sql.NullString `json:"ref6"`
	Ref7                 sql.NullString `json:"ref7"`
}
