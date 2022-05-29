package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var productStore = make(map[string]Product)
var id int = 0

func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)
	id++
	product.ID = id
	product.CreatedOn = time.Now()
	product.ChangedOn = time.Now()
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product

	for _, v := range productStore {
		products = append(products, v)
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})

	data, err := json.Marshal(products)
	CheckError(err)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	vars := mux.Vars(r)

	key, err := strconv.Atoi(vars["id"])
	CheckError(err)

	for _, v := range productStore {
		if v.ID == key {
			product = v
		}
	}

	data, err := json.Marshal(product)
	CheckError(err)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	var updProduct Product

	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		key, err := strconv.Atoi(vars["id"])
		CheckError(err)
		for _, v := range productStore {
			if v.ID == key {
				updProduct.ID = v.ID
				updProduct.CreatedOn = v.CreatedOn
				delete(productStore, strconv.Itoa(v.ID))
			}
		}
	} else {
		fmt.Println("Bu ID'ye sahip ürün bulunamadı.")
	}

	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)

	updProduct.Name = product.Name
	updProduct.Description = product.Description
	updProduct.ChangedOn = time.Now()

	productStore[key] = updProduct

	data, err := json.Marshal(updProduct)
	CheckError(err)
	w.Header().Set("Convert-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	//var product Product

	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	CheckError(err)

	for _, v := range productStore {
		if v.ID == key {
			delete(productStore, strconv.Itoa(key))
		}
	}

	w.WriteHeader(http.StatusOK)

}

type Product struct {
	ID          int       `json :"id"`
	Name        string    `json :"name"`
	Description string    `json : "description"`
	CreatedOn   time.Time `json : "createdon"`
	ChangedOn   time.Time `json : "changedon`
}

func CheckError(err error) {
	if err != nil {
		error.Error(err)
	}
}
