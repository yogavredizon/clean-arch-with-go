package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yogavredizon/clean-arch/model"
	"github.com/yogavredizon/clean-arch/service"
)

type ProductHandler struct {
	ProductService service.ProductServiceInterface
}

func (p *ProductHandler) CreateProductHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		var err error

		err = json.NewDecoder(r.Body).Decode(&product)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		product, err = p.ProductService.Create(product)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := model.Response{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
			Data:    product,
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}

func (p *ProductHandler) FindProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		id := r.URL.Query().Get("id")

		product, err := p.ProductService.FindById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := model.Response{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    product,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
