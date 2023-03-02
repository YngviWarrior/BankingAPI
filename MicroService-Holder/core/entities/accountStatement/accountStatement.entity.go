package accountStatementEntity

type AccountStatement struct {
	AccountStatement uint64  `json:"account_statement"`
	Account          uint64  `json:"account"`
	TransactionType  uint64  `json:"transaction_type"`
	PreviousBalance  float64 `json:"previous_balance"`
	CurrentBalance   float64 `json:"current_balance"`
	RegisteredDate   string  `json:"registered_at"`
}

type AccountStatementTransactionType struct {
	AccountStatement           uint64  `json:"account_statement"`
	Account                    uint64  `json:"account"`
	TransactionType            uint64  `json:"transaction_type"`
	TransactionTypeDescription string  `json:"transaction_type_description"`
	PreviousBalance            float64 `json:"previous_balance"`
	CurrentBalance             float64 `json:"current_balance"`
	RegisteredDate             string  `json:"registered_at"`
}
