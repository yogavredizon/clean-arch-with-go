package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogavredizon/clean-arch/model"
)

// setup repository mock
type RepositoryMock struct {
	mock.Mock
}

func (rm *RepositoryMock) Create(p model.Product) (model.Product, error) {
	args := rm.Called(p)

	if args.Get(1) != nil {
		return model.Product{}, errors.New("failed to add product")
	}

	product := args.Get(0).(model.Product)
	return product, nil
}

func (rm *RepositoryMock) FindById(id string) (model.Product, error) {
	args := rm.Called(id)

	if args.Get(1) != nil {
		if id == "" {
			return model.Product{}, errors.New("id can't empty")
		}

		return model.Product{}, errors.New("id not found")
	}

	product := args.Get(0).(model.Product)
	return product, nil
}

var repo = new(RepositoryMock)
var service = NewProductService(repo)

var value = model.Product{
	Id:          "1",
	Name:        "Laptop",
	Description: "Laptop Core I5",
	Location:    "Jakarta",
	Price:       10000000,
	Sales:       nil,
}

func Test_Service_To_Create_NewProduct(t *testing.T) {
	repo.On("Create", value).Return(value, nil)

	product, err := service.Create(value)
	assert.Equal(t, value, product)
	assert.Nil(t, err)
}
func Test_Service_To_Create_NewProduct_With_Empty_ID(t *testing.T) {
	// set id to be empty string
	value.Id = ""
	repo.On("Create", value).Return(model.Product{}, errors.New("id can't empty"))

	product, err := service.Create(value)
	assert.Equal(t, model.Product{}, product)
	assert.NotNil(t, err)

	// reset id to default
	value.Id = "1"
}
func Test_Service_To_Find_Product_With_Empty_ID(t *testing.T) {
	repo.On("FindById", "").Return(model.Product{}, errors.New("id can't empty"))

	product, err := service.FindById("")
	assert.Equal(t, model.Product{}, product)
	assert.NotNil(t, err)
}
func Test_Service_To_Find_Product_With_Exits_ID(t *testing.T) {
	repo.On("FindById", "1").Return(value, nil)

	product, err := service.FindById("1")
	assert.Equal(t, value, product)
	assert.Nil(t, err)
}
