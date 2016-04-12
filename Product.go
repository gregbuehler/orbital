package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	Name        string `json:"name,required" description:"The name of the product"`
	Description string `json:"description" description:"The description of the product"`
	ProductID   string `json:"product_id,required" description:"The user-visible id of the product"`
	Secret      string `json:"secret,required" description:"The product hex string that all devices belonging to this product will use to generate activation codes"`
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductSimpleListHandler!\n\n%s", r.RequestURI)
}

func ProductCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductCreateHandler!\n\n%s", r.RequestURI)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductDetailHandler!\n\n%s", r.RequestURI)
}

func ProductUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductUpdateHandler!\n\n%s", r.RequestURI)
}

func ProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductDeleteHandler!\n\n%s", r.RequestURI)
}
