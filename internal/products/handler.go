package products

import (
	"MarketPlace/internal/middleware"
	"encoding/json"
	"net/http"
	"strconv"
)

type Product struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	UserID int     `json:"user_id"`
}

func GetProduct(service *ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := service.GetProduct(r.Context())
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func AddProduct(service *ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product
		ctx := r.Context()
		userID := ctx.Value(middleware.UserIDkey)

		value, ok := userID.(int)

		if !ok {
			http.Error(w, "error", http.StatusUnauthorized)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.CreateProduct(r.Context(), product.Name, product.Price, value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	}
}

func DeleteProduct(service *ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value(middleware.UserIDkey)
		value, ok := userID.(int)
		if !ok {
			http.Error(w, "error", http.StatusUnauthorized)
			return
		}
		productId := r.PathValue("id")
		productIdInt, err := strconv.Atoi(productId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = service.DeleteProduct(r.Context(), productIdInt, value)
		if err != nil {
			http.Error(w, "not auth", http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
