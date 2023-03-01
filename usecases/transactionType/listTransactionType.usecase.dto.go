package transactionTypeUseCase

type InputListTransactionTypeDto struct{}

type transactionType struct {
	TransactionType uint64 `json:"transaction_type"`
	Description     string `json:"description"`
}

type OutputListTransactionTypeDto struct {
	List []*transactionType `json:"types_list,omitempty"`
}
