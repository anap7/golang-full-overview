package main

import (
	"github.com/anap7/golang-overviews/src/http"
	"github.com/anap7/golang-overviews/src/model"
)

func main() {
	p1 := model.NewProduct()
	p1.Name = "Carrinho"

	p2 := model.NewProduct()
	p2.Name = "Boneca"

	products := model.Products{}
	products.Add(*p1)
	products.Add(*p2)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()
}
