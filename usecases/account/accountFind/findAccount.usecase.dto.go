package findAccountUseCase

type InputFindAccountDto struct {
	Agency string
	Number string
}

type OutputFindAccountDto struct {
	HolderName string  `json:"holder_name"`
	HolderDoc  string  `json:"holder_doc"`
	Agency     string  `json:"agency"`
	Number     string  `json:"number"`
	Balance    float64 `json:"balance"`
	Activated  bool    `json:"activated"`
	Blocked    bool    `json:"blocked"`
}
