package products

import (
	"MarketPlace/internal/User"
	"context"
	"errors"
	"unicode/utf8"
)

type ProductService struct {
	repo        *ProductRepository
	userService *User.UserService
}

func NewProductService(repo *ProductRepository, userService *User.UserService) *ProductService {
	return &ProductService{repo: repo, userService: userService}
}

func (s *ProductService) GetProduct(ctx context.Context) ([]Product, error) {
	return s.repo.GetProduct(ctx)
}

func (s *ProductService) CreateProduct(ctx context.Context, name string, price float64, value int) error {
	if utf8.RuneCountInString(name) > 20 {
		return errors.New("Character limit exceeded")
	}
	if price <= 0.9 {
		return errors.New("Price is too low")
	}
	role, err := s.userService.IsSeller(ctx, value)
	if err != nil {
		errors.New("No seller")
		return err
	}
	if role != "seller" {
		return errors.New("No create buyer")

	}
	err = s.repo.ProductCreate(ctx, name, price, value)
	if err != nil {
		errors.New("DB Errors")
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, product_id int, value int) error {
	role, err := s.userService.IsSeller(ctx, value)
	if err != nil {
		return err
	}
	if role != "seller" && role != "moderator" {
		return errors.New("No seller or moderator")
	}
	count, err := s.repo.DeleteProduct(ctx, product_id, value, role)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("No delete")
	}
	return nil

}
