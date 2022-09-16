package validator

type Validator struct {
	Error []string
}

type ValidatorInterface interface {
	InputValidator(o any) []string
}
