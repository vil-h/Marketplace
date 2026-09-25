package main

import (
	"MarketPlace/internal/User"
	"MarketPlace/internal/database"
	"MarketPlace/internal/middleware"
	"MarketPlace/internal/products"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	db, err := database.NewPostgres(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	jwtSec := os.Getenv("JWT_SECRET")
	repo := products.NewProductRepository(db)
	repoUser := User.NewUserRepository(db)
	serviceUser := User.NewUserService(repoUser, jwtSec)
	service := products.NewProductService(repo, serviceUser)
	mux := http.NewServeMux()
	mux.Handle("GET /me", middleware.Auth(jwtSec)(User.Me(serviceUser)))
	mux.Handle("POST /login", User.Login(serviceUser))
	mux.Handle("GET /product", products.GetProduct(service))
	mux.Handle("POST /register", User.CreateUser(serviceUser))
	mux.Handle("POST /createproduct", middleware.Auth(jwtSec)(products.AddProduct(service)))
	mux.Handle("DELETE /product/{id}", middleware.Auth(jwtSec)(products.DeleteProduct(service)))
	err = http.ListenAndServe(":8080", middleware.Cors(mux))
	if err != nil {
		log.Fatal(err)
	}
}
