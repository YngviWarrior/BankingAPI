package transactionAccountUseCase

type InputTransactionAccountDto struct {
	Agency          string
	Number          string
	TransactionType uint64
	Amount          float64
}

type OutputTransactionAccountDto struct {
}
