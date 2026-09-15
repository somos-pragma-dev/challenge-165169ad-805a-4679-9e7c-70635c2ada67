package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null;uniqueIndex" json:"name"`
	Price     float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock     int            `gorm:"type:integer;not null;default:0" json:"stock"`
	Category  string         `gorm:"type:varchar(100);not null" json:"category"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return &ValidationError{Field: "name", Message: "el nombre del producto es requerido"}
	}
	if len(p.Name) > 255 {
		return &ValidationError{Field: "name", Message: "el nombre no puede exceder 255 caracteres"}
	}
	if p.Price < 0 {
		return &ValidationError{Field: "price", Message: "el precio no puede ser negativo"}
	}
	if p.Stock < 0 {
		return &ValidationError{Field: "stock", Message: "el stock no puede ser negativo"}
	}
	if p.Category == "" {
		return &ValidationError{Field: "category", Message: "la categoria es requerida"}
	}
	if len(p.Category) > 100 {
		return &ValidationError{Field: "category", Message: "la categoria no puede exceder 100 caracteres"}
	}
	return nil
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) GetField() string {
	return e.Field
}