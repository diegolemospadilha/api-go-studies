package usecase

import "github.com/diegolemospadilha/api-products-go/internal/entity"

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*ListProductsOutputDto

	for _, product := range products {
		output = append(output, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return output, err

}
