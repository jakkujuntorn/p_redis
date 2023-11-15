package russyservices

type IRus_service interface {
	Custom_GetProduct() ([]dataPoductService, error)
}

type dataPoductService struct {
	ID       int
	Name     string
	Quantity int
}
