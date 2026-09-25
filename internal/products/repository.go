package products

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProduct(ctx context.Context) ([]Product, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, name, price, user_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.UserID,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) ProductCreate(ctx context.Context, name string, price float64, value int) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO products(name, price, user_id)
				VALUES( $1, $2, $3)`, name, price, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, product_id int, value int, role string) (int64, error) {
	result, err := r.db.Exec(
		ctx,
		`DELETE FROM products
		WHERE id = $1
		AND (user_id = $2
		OR $3 = 'moderator');`, product_id, value, role)
	if err != nil {
		return 0, err
	}
	count := result.RowsAffected()
	return count, err
}
