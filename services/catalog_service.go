package services

// ****** Adaptor ******

import (
	"fmt"
	"goredis/repositories"
)

type catalogService struct {
	test string
	lastName string
	productRepo repositories.I_ProducrRepository
}

func NewCatalogService(product_Repo repositories.I_ProducrRepository) I_CatalogService {

	return catalogService{productRepo: product_Repo}
}

func (s catalogService) GetProducts() (products []ProductCatalog, err error) {

	// จะ return ออกไปเลยก็ได้
	// productRepo.GetProducts() มาจาก layer repositories ***
	products_DB, err := s.productRepo.GetProducts()

	if err != nil {
		// ตรงนี้จะ custom error ด้วยก็ได้
		return nil, err
	}

	// ปั้นข้อมูลใหม่ เพื่อไม่ให้เกี่ยวข้องกับ struct product ของ repo แค่นั้น
	// ตัดขาดระหว่าง DB (repo) กับ service
	for _, p := range products_DB {
		products = append(products, ProductCatalog{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	fmt.Println(" ******* Data base Service****** ")
	return products, nil
}
