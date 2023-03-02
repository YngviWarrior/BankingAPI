package blockAccountUseCase

type InputBlockAccountDto struct {
	Agency string
	Number string
	Block  bool
}

type OutputBlockAccountDto struct{}
