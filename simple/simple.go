package simple

type SimpleRepository struct {
}

// simpleservice bergantung pada repository
type SimpleService struct {
	*SimpleRepository
}

// provider (biasanya constructor nya, pake data New depan nya)
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// provider (biasanya constructor nya, pake data New depan nya)
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
