package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yogavredizon/clean-arch/model"
)

var repo = NewProductsRepository()

func Test_Create_Product_When_Input_is_Wrong(t *testing.T) {
	product, err := repo.Create(model.Product{})

	assert.Equal(t, model.Product{}, product)
	assert.NotNil(t, err)
}

func Test_Create_Product_When_Input_is_Rigth(t *testing.T) {
	value := model.Product{
		Id:          "1",
		Name:        "Laptop",
		Description: "Laptop Core I5",
		Location:    "Jakarta",
		Price:       10000000,
		Sales:       nil,
	}
	product, err := repo.Create(value)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(model.Products))
	assert.Equal(t, value, product)
}

func Test_FindByID(t *testing.T) {
	model.Products = []model.Product{
		{
			Id:          "1",
			Name:        "Laptop",
			Description: "Laptop Core I5",
			Location:    "Jakarta",
			Price:       10000000,
			Sales:       nil,
		},
		{
			Id:          "2",
			Name:        "Laptop",
			Description: "Laptop Core I5",
			Location:    "Jakarta",
			Price:       10000000,
			Sales:       nil,
		},
	}

	testCase := []struct {
		Title string
		Data  string
		Error error
	}{
		{Title: "When Id is empty", Data: "", Error: errors.New("id can't empty")},
		{Title: "When Id not in Instance DB", Data: "9", Error: errors.New("id not found")},
		{Title: "When Id is exists", Data: "1", Error: nil},
	}

	for _, tt := range testCase {
		t.Run(tt.Title, func(t *testing.T) {
			_, err := repo.FindById(tt.Data)
			assert.Equal(t, err, tt.Error)
		})
	}
}
