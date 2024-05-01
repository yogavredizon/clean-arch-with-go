package repository

import (
	"errors"

	"github.com/yogavredizon/clean-arch/model"
)

type ProductInterface interface {
	Create(product model.Product) (model.Product, error)
	FindById(id string) (model.Product, error)
}

type ProductImpl struct {
}

func NewProductsRepository() ProductInterface {
	return &ProductImpl{}
}

func (rp *ProductImpl) Create(product model.Product) (model.Product, error) {
	// get first length of product
	s := len(model.Products)

	// add product to instance DB
	model.Products = append(model.Products, product)

	// get length after add product to instance DB
	l := len(model.Products)

	// check if the last length is lower than the first length, it will send error
	if s >= l {
		return model.Product{}, errors.New("failed to add product")
	}
	return product, nil
}

func (rp *ProductImpl) FindById(id string) (model.Product, error) {
	if id == "" {
		return model.Product{}, errors.New("id can't empty")
	}

	for _, product := range model.Products {
		if product.Id == id {
			return product, nil
		}
	}

	return model.Product{}, errors.New("id not found")
}
