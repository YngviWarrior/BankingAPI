package listStatementUseCase

type InputListStatementDto struct {
	Account    uint64
	DateStart  string
	DateFinish string
}

type statementList struct {
	TransactionTypeDescription string  `json:"transaction_type_description"`
	PreviousBalance            float64 `json:"previous_balance"`
	CurrentBalance             float64 `json:"current_balance"`
	RegisteredDate             string  `json:"registered_at"`
}

type OutputListStatementDto struct {
	List []*statementList `json:"statements,omitempty"`
}
