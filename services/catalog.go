package services

// ****** plug ******

// Product เอาไว้คั่นกลางระหว่าง repository กับ handlers
type ProductCatalog struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type I_CatalogService interface {
	GetProducts() ([]ProductCatalog, error)
}
