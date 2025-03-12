package product

type ProductUsecaseProblem struct {
	repository *ProductRepository
}
type ProductUsecase struct {
	repository ProductRepositoryInterface
}

func NewProductUsecaseProblem(repository *ProductRepository) *ProductUsecaseProblem {
	return &ProductUsecaseProblem{
		repository: repository,
	}
}
func NewProductUsecase(repository ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

// This Product was not supossed to be returned. We should return a DTO instead.
// However, we will keep it for the sake of the example
func (p *ProductUsecaseProblem) GetProductProblem(id int) (Product, error) {
	return p.repository.GetProduct(id)
}
func (p *ProductUsecase) GetProduct(id int) (Product, error) {
	return p.repository.GetProduct(id)
}
