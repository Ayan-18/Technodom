package repository

import (
	"database/sql"
	"errors"
	"log/slog"
	"technodom/models"
)

type ProductRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewProductRepository(db *sql.DB, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{
		db:     db,
		logger: logger,
	}
}

// CreateProduct добавляет новый продукт в базу данных.
func (r *ProductRepository) CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, category, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	r.logger.Info("Creating product", slog.String("name", product.Name))

	err := r.db.QueryRow(query, product.Name, product.Description, product.Price,
		product.Category, product.CreatedAt, product.UpdatedAt).Scan(&product.ID)
	if err != nil {
		r.logger.Error("Error creating product", slog.String("error", err.Error()))
		return err
	}

	r.logger.Info("Product created", slog.Int("id", product.ID))
	return nil
}

// GetProductByID получает продукт по ID.
func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := `SELECT id, name, description, price, category, created_at, updated_at FROM products WHERE id = $1`
	r.logger.Info("Getting product by ID", slog.Int("id", id))

	product := &models.Product{}
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description,
		&product.Price, &product.Category, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Warn("Product not found", slog.Int("id", id))
			return nil, nil
		}
		r.logger.Error("Error retrieving product", slog.String("error", err.Error()))
		return nil, err
	}

	return product, nil
}

// UpdateProduct обновляет информацию о продукте.
func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, category = $4, updated_at = $5 WHERE id = $6`
	r.logger.Info("Updating product", slog.Int("id", product.ID))

	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Category, product.UpdatedAt, product.ID)
	if err != nil {
		r.logger.Error("Error updating product", slog.String("error", err.Error()))
		return err
	}

	r.logger.Info("Product updated", slog.Int("id", product.ID))
	return nil
}

// DeleteProduct удаляет продукт по ID.
func (r *ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	r.logger.Info("Deleting product", slog.Int("id", id))

	_, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Error("Error deleting product", slog.String("error", err.Error()))
		return err
	}

	r.logger.Info("Product deleted", slog.Int("id", id))
	return nil
}

// GetAllProducts возвращает все продукты из базы данных.
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := `SELECT id, name, description, price, category, created_at, updated_at FROM products`
	r.logger.Info("Getting all products")

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error("Error retrieving products", slog.String("error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price,
			&product.Category, &product.CreatedAt, &product.UpdatedAt); err != nil {
			r.logger.Error("Error scanning product", slog.String("error", err.Error()))
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error iterating rows", slog.String("error", err.Error()))
		return nil, err
	}

	r.logger.Info("Retrieved products", slog.Int("count", len(products)))
	return products, nil
}
