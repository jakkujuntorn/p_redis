package russyservices

import (
	"goredis/repositories"
)

type ProductServices struct {
	productRepo repositories.I_ProducrRepository
}

func NewR_Catalog_Service(proREpo repositories.I_ProducrRepository) IRus_service {
	return ProductServices{proREpo}
}

func (s ProductServices) Custom_GetProduct() (products []dataPoductService, err error) {

	data, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range data {

		products = append(products, dataPoductService{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})

	}

	return products, nil

}
