package listStatementUseCase

type InputListStatementDto struct {
	Account    uint64
	DateStart  string
	DateFinish string
}

type OutputListStatementDto struct {
	List []struct {
		AccountStatement uint64  `json:"account_statement"`
		Account          uint64  `json:"account"`
		TransactionType  uint64  `json:"transaction_type"`
		PreviousBalance  float64 `json:"previous_balance"`
		CurrentBalance   float64 `json:"current_balance"`
		RegisteredDate   string  `json:"registered_at"`
	} `json:"list"`
}
