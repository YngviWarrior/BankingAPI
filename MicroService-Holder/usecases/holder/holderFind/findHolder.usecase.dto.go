package findHolderUseCase

type InputFindHolderDto struct {
	CPF string
}

type account struct {
	HolderName string  `json:"holder_name"`
	HolderDoc  string  `json:"holder_doc"`
	Agency     string  `json:"agency"`
	Number     string  `json:"number"`
	Balance    float64 `json:"balance"`
	Activated  bool    `json:"activated"`
	Blocked    bool    `json:"blocked"`
}

type OutputFindHolderDto struct {
	FullName    string     `json:"full_name"`
	CPF         string     `json:"cpf"`
	Verified    bool       `json:"verified"`
	Activated   bool       `json:"activated"`
	AccountList []*account `json:"account_list,omitempty"`
}
