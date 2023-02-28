package holderEntity

type Holder struct {
	Holder    uint64 `json:"holder"`
	FullName  string `json:"full_name"`
	CPF       string `json:"cpf"`
	Verified  bool   `json:"verified"`
	Activated bool   `json:"activated"`
}
