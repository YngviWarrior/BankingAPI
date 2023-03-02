package database

type MockRepository struct{}

type MockRepositoryInterface interface {
	Find() (s string)
	FindAll() (list []string)
	Create() bool
	Update() bool
}

func (MockRepository) Find() (s string) {
	return "find test"
}

func (MockRepository) FindAll() (list []string) {
	list = append(list, "test")
	list = append(list, "test1")

	return
}

func (MockRepository) Create() bool {
	return true
}

func (MockRepository) Update() bool {
	return true
}
