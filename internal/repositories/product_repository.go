package repositories

import (
	"errors"
	"fmt"

	"github.com/empresa/catalogo-productos/internal/models"
	"gorm.io/gorm"
)

var (
	ErrProductNotFound = errors.New("producto no encontrado")
	ErrDuplicateProduct = errors.New("ya existe un producto con ese nombre")
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll() ([]models.Product, error)
	FindByName(name string) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
	ExistsByName(name string) (bool, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *models.Product) error {
	exists, err := r.ExistsByName(product.Name)
	if err != nil {
		return fmt.Errorf("error al verificar existencia del producto: %w", err)
	}
	if exists {
		return ErrDuplicateProduct
	}

	result := r.db.Create(product)
	if result.Error != nil {
		return fmt.Errorf("error al crear el producto en la base de datos: %w", result.Error)
	}
	return nil
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("error al buscar el producto por ID: %w", result.Error)
	}
	return &product, nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("error al listar los productos: %w", result.Error)
	}
	return products, nil
}

func (r *productRepository) FindByName(name string) (*models.Product, error) {
	var product models.Product
	result := r.db.Where("name = ?", name).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("error al buscar el producto por nombre: %w", result.Error)
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	result := r.db.Save(product)
	if result.Error != nil {
		return fmt.Errorf("error al actualizar el producto: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func (r *productRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("error al eliminar el producto: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func (r *productRepository) ExistsByName(name string) (bool, error) {
	var count int64
	result := r.db.Model(&models.Product{}).Where("name = ?", name).Count(&count)
	if result.Error != nil {
		return false, fmt.Errorf("error al verificar existencia del producto por nombre: %w", result.Error)
	}
	return count > 0, nil
}