-- Migration: Create products table
-- Description: Crea la tabla de productos con las restricciones necesarias
-- Author: Sistema de Migraciones
-- Created: 2024

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Índice único para el nombre del producto
CREATE UNIQUE INDEX IF NOT EXISTS idx_products_name_unique ON products (LOWER(name));

-- Índice para búsquedas por categoría
CREATE INDEX IF NOT EXISTS idx_products_category ON products (category);

-- Índice para búsquedas por nombre (búsqueda parcial)
CREATE INDEX IF NOT EXISTS idx_products_name_search ON products (name);

-- Constraint: precio no puede ser negativo
ALTER TABLE products
ADD CONSTRAINT chk_products_price_positive CHECK (price >= 0);

-- Constraint: stock no puede ser negativo
ALTER TABLE products
ADD CONSTRAINT chk_products_stock_non_negative CHECK (stock >= 0);

-- Constraint: nombre no puede estar vacío
ALTER TABLE products
ADD CONSTRAINT chk_products_name_not_empty CHECK (LENGTH(TRIM(name)) > 0);

-- Constraint: categoría no puede estar vacía
ALTER TABLE products
ADD CONSTRAINT chk_products_category_not_empty CHECK (LENGTH(TRIM(category)) > 0);

-- Comentarios para documentación
COMMENT ON TABLE products IS 'Tabla principal para el catálogo de productos';
COMMENT ON COLUMN products.name IS 'Nombre del producto, debe ser único';
COMMENT ON COLUMN products.price IS 'Precio del producto en formato decimal';
COMMENT ON COLUMN products.stock IS 'Cantidad disponible en inventario';
COMMENT ON COLUMN products.category IS 'Categoría a la que pertenece el producto';