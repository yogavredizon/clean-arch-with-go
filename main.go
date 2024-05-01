package main

import (
	"net/http"

	"github.com/yogavredizon/clean-arch/controller"
	"github.com/yogavredizon/clean-arch/middleware"
	"github.com/yogavredizon/clean-arch/repository"
	"github.com/yogavredizon/clean-arch/service"
)

func main() {
	repo := repository.NewProductsRepository()
	service := service.NewProductService(repo)
	handler := controller.ProductHandler{
		ProductService: service,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/products", middleware.RequestPost(handler.CreateProductHandler()))
	mux.HandleFunc("/api/v1/product", middleware.RequestGet(handler.FindProductByIdHandler()))

	http.ListenAndServe(":8080", mux)
}
