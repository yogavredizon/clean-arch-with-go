package service

import (
	"errors"

	"github.com/yogavredizon/clean-arch/model"
	"github.com/yogavredizon/clean-arch/repository"
)

type ProductServiceInterface interface {
	Create(product model.Product) (model.Product, error)
	FindById(id string) (model.Product, error)
}

type ProductService struct {
	Repository repository.ProductInterface
}

func NewProductService(repository repository.ProductInterface) ProductServiceInterface {
	return &ProductService{repository}
}

func (ps *ProductService) Create(product model.Product) (model.Product, error) {
	if product.Id == "" {
		return model.Product{}, errors.New("id can't empty")
	}
	return ps.Repository.Create(product)
}
func (ps *ProductService) FindById(id string) (model.Product, error) {
	return ps.Repository.FindById(id)
}
