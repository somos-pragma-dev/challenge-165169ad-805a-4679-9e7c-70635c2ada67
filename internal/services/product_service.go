package services

import (
	"errors"

	"github.com/empresa/catalogo-productos/internal/dto"
	"github.com/empresa/catalogo-productos/internal/models"
	"github.com/empresa/catalogo-productos/internal/repositories"
	"github.com/empresa/catalogo-productos/internal/utils"
)

var (
	ErrProductNotFound     = errors.New("producto no encontrado")
	ErrProductAlreadyExists = errors.New("ya existe un producto con ese nombre")
	ErrInvalidPrice        = errors.New("el precio no puede ser negativo")
	ErrInvalidStock        = errors.New("el stock no puede ser negativo")
	ErrInvalidName         = errors.New("el nombre no puede estar vacío")
)

type ProductService interface {
	CreateProduct(req *dto.CreateProductRequest) (*models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	UpdateProduct(id uint, req *dto.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(pr repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: pr,
	}
}

func (s *productService) CreateProduct(req *dto.CreateProductRequest) (*models.Product, error) {
	if err := utils.ValidateCreateProductRequest(req); err != nil {
		return nil, err
	}

	exists, err := s.productRepository.ExistsByName(req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &models.ValidationError{
			Field:   "name",
			Message: ErrProductAlreadyExists.Error(),
		}
	}

	product := &models.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		Category: req.Category,
	}

	if err := s.productRepository.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

func (s *productService) ListProducts() ([]models.Product, error) {
	return s.productRepository.FindAll()
}

func (s *productService) UpdateProduct(id uint, req *dto.UpdateProductRequest) (*models.Product, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, ErrProductNotFound
	}

	if err := utils.ValidateUpdateProductRequest(req); err != nil {
		return nil, err
	}

	if req.Name != nil && *req.Name != product.Name {
		exists, err := s.productRepository.ExistsByName(*req.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, &models.ValidationError{
				Field:   "name",
				Message: ErrProductAlreadyExists.Error(),
			}
		}
		product.Name = *req.Name
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	if req.Category != nil {
		product.Category = *req.Category
	}

	if err := s.productRepository.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id uint) error {
	exists, err := s.productRepository.FindByID(id)
	if err != nil {
		return ErrProductNotFound
	}
	if exists == nil {
		return ErrProductNotFound
	}

	return s.productRepository.Delete(id)
}