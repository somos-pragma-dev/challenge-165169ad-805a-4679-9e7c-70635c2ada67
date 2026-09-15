package dto

import (
	"time"
)

type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required,max=255"`
	Price    float64 `json:"price" binding:"required,min=0"`
	Stock    int     `json:"stock" binding:"required,min=0"`
	Category string  `json:"category" binding:"required,max=100"`
}

type UpdateProductRequest struct {
	Name     *string  `json:"name" binding:"omitempty,max=255"`
	Price    *float64 `json:"price" binding:"omitempty,min=0"`
	Stock    *int     `json:"stock" binding:"omitempty,min=0"`
	Category *string  `json:"category" binding:"omitempty,max=100"`
}

type ProductResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int64             `json:"total"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ToProductResponse(name string, price float64, stock int, category string, id uint, createdAt, updatedAt time.Time) ProductResponse {
	return ProductResponse{
		ID:        id,
		Name:      name,
		Price:     price,
		Stock:     stock,
		Category:  category,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToProductResponseFromModel(model interface{}) ProductResponse {
	switch m := model.(type) {
	case *struct {
		ID        uint
		Name      string
		Price     float64
		Stock     int
		Category  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}:
		return ProductResponse{
			ID:        m.ID,
			Name:      m.Name,
			Price:     m.Price,
			Stock:     m.Stock,
			Category:  m.Category,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		}
	default:
		return ProductResponse{}
	}
}