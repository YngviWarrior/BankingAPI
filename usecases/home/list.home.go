package homeUseCase

func (h *HomeUseCase) ListAll() (output []string) {
	output = h.HomeRepository.FindAll()

	return
}
