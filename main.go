package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/products", ProductHandle)
	http.ListenAndServe(":9095", nil)
}

func ProductHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola jovem"))
}

/*func main() {
	p1 := model.NewProduct()
	p1.Name = "Carrinho"

	p2 := model.NewProduct()
	p2.Name = "Boneca"

	products := model.Products{}
	products.Add(*p1)
	products.Add(*p2)
}*/
