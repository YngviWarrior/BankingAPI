package passrecovery

type InputPassRecoveryDto struct {
	Email string `json:"email"`
}

type OutputPassRecoveryDto struct {
	InternalStatus int64
}
