package accountEntity

type Account struct {
	Account   uint64  `json:"account"`
	Holder    uint64  `json:"holder"`
	Agency    string  `json:"agency"`
	Number    string  `json:"number"`
	Balance   float64 `json:"balance"`
	Activated bool    `json:"activated"`
	Blocked   bool    `json:"blocked"`
}

type AccountHolder struct {
	Account         uint64  `json:"account"`
	Holder          uint64  `json:"holder"`
	Agency          string  `json:"agency"`
	Number          string  `json:"number"`
	Balance         float64 `json:"balance"`
	Activated       bool    `json:"activated"`
	Blocked         bool    `json:"blocked"`
	HolderName      string  `json:"full_name"`
	HolderDoc       string  `json:"cpf"`
	HolderVerified  bool    `json:"verified"`
	HolderActivated bool    `json:"holder_activated"`
}
