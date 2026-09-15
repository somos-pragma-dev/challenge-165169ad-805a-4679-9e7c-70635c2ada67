package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/empresa/catalogo-productos/internal/dto"
	"github.com/empresa/catalogo-productos/internal/models"
)

var (
	ErrEmptyName     = errors.New("el nombre no puede estar vacío")
	ErrNegativePrice = errors.New("el precio no puede ser negativo")
	ErrNegativeStock = errors.New("el stock no puede ser negativo")
	ErrInvalidCategory = errors.New("la categoría no puede estar vacía")
)

func ValidateCreateProductRequest(req *dto.CreateProductRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return &models.ValidationError{
			Field:   "name",
			Message: ErrEmptyName.Error(),
		}
	}

	if req.Price < 0 {
		return &models.ValidationError{
			Field:   "price",
			Message: ErrNegativePrice.Error(),
		}
	}

	if req.Stock < 0 {
		return &models.ValidationError{
			Field:   "stock",
			Message: ErrNegativeStock.Error(),
		}
	}

	if strings.TrimSpace(req.Category) == "" {
		return &models.ValidationError{
			Field:   "category",
			Message: ErrInvalidCategory.Error(),
		}
	}

	return nil
}

func ValidateUpdateProductRequest(req *dto.UpdateProductRequest) error {
	if req.Name != nil {
		if strings.TrimSpace(*req.Name) == "" {
			return &models.ValidationError{
				Field:   "name",
				Message: ErrEmptyName.Error(),
			}
		}
	}

	if req.Price != nil {
		if *req.Price < 0 {
			return &models.ValidationError{
				Field:   "price",
				Message: ErrNegativePrice.Error(),
			}
		}
	}

	if req.Stock != nil {
		if *req.Stock < 0 {
			return &models.ValidationError{
				Field:   "stock",
				Message: ErrNegativeStock.Error(),
			}
		}
	}

	if req.Category != nil {
		if strings.TrimSpace(*req.Category) == "" {
			return &models.ValidationError{
				Field:   "category",
				Message: ErrInvalidCategory.Error(),
			}
		}
	}

	return nil
}

func ValidateProductName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("el nombre del producto no puede estar vacío")
	}
	if len(name) < 3 {
		return fmt.Errorf("el nombre del producto debe tener al menos 3 caracteres")
	}
	if len(name) > 100 {
		return fmt.Errorf("el nombre del producto no puede exceder 100 caracteres")
	}
	return nil
}

func ValidatePrice(price float64) error {
	if price < 0 {
		return ErrNegativePrice
	}
	if price > 999999.99 {
		return fmt.Errorf("el precio excede el límite permitido")
	}
	return nil
}

func ValidateStock(stock int) error {
	if stock < 0 {
		return ErrNegativeStock
	}
	if stock > 999999 {
		return fmt.Errorf("el stock excede el límite permitido")
	}
	return nil
}