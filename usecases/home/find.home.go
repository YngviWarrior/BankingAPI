package homeUseCase

func (h *HomeUseCase) Find() (output string) {
	output = h.HomeRepository.Find()

	return
}
