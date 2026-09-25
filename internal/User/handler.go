package User

import (
	"MarketPlace/internal/middleware"
	"encoding/json"
	"net/http"
	"time"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	Role         string    `json:"role"`
}

func CreateUser(service *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		err = service.Register(r.Context(), req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func Login(service *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tok, err := service.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tok)
	}
}

func Me(service *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UserID := r.Context().Value(middleware.UserIDkey)
		userId, ok := UserID.(int)
		if !ok {
			http.Error(w, "error", http.StatusUnauthorized)
			return
		}
		role, email, err := service.Me(r.Context(), userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		users := User{Role: role, Email: email}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}

}
