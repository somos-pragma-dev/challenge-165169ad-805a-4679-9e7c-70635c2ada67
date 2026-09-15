# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener un ZIP con el proyecto completo y arrancable.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Como saber que terminaste

```bash
el comando de build o arranque canonico del stack elegido
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
API REST con Go, Gin framework y GORM

### Reto
- Tema: Go Gin API
- Seniority: junior-l2
- Tipo: practical
- Título: Implementación de una API REST en Go con Gin y GORM
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Registro de productos — objetivo: Implementar la funcionalidad para registrar productos en la API. — entregable (NO resolver): API que acepta solicitudes POST para registrar productos con validación de precio y nombre.
- Fase 2: Listado de productos — objetivo: Implementar la funcionalidad para listar productos en la API. — entregable (NO resolver): API que acepta solicitudes GET para listar productos.
- Fase 3: Actualización de productos — objetivo: Implementar la funcionalidad para actualizar productos en la API. — entregable (NO resolver): API que acepta solicitudes PUT para actualizar productos con validación de precio y nombre.
- Fase 4: Eliminación de productos — objetivo: Implementar la funcionalidad para eliminar productos en la API. — entregable (NO resolver): API que acepta solicitudes DELETE para eliminar productos.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/empresa/catalogo-productos

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/go-playground/validator/v10 v10.20.0
	gorm.io/driver/postgres v1.5.5
	gorm.io/gorm v1.25.7
)

// Indirect dependencies for type clarity in go.mod
require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.4 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// === ARCHIVO: cmd/api/main.go ===
package main

import (
	"log"
	"os"

	"github.com/empresa/catalogo-productos/config"
	"github.com/empresa/catalogo-productos/internal/handlers"
	"github.com/empresa/catalogo-productos/internal/models"
	"github.com/empresa/catalogo-productos/internal/repositories"
	"github.com/empresa/catalogo-productos/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Error al ejecutar AutoMigrate: %v", err)
	}

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "API de catálogo de productos funcionando"})
	})

	api := router.Group("/api/v1")
	{
		api.POST("/products", productHandler.CreateProduct)
		api.GET("/products", productHandler.ListProducts)
		api.GET("/products/:id", productHandler.GetProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando en el puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// === ARCHIVO: internal/repositories/product_repository.go ===
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

// === ARCHIVO: internal/models/product.go ===
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

// === ARCHIVO: internal/dto/product_dto.go ===
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
	case struct {
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

// === ARCHIVO: config/database.go ===
package config

import (
	"fmt"
	"os"

	"github.com/empresa/catalogo-productos/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "catalogo_productos"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func InitDatabase() (*gorm.DB, error) {
	config := getDatabaseConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener la conexion de la base de datos: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

func GetDSN() string {
	config := getDatabaseConfig()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)
}


// === ARCHIVO: internal/handlers/product_handler.go ===
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/empresa/catalogo-productos/internal/dto"
	"github.com/empresa/catalogo-productos/internal/models"
	"github.com/empresa/catalogo-productos/internal/services"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(ps services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: ps,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Error de validación",
			Message: err.Error(),
		})
		return
	}

	product, err := h.productService.CreateProduct(&req)
	if err != nil {
		switch e := err.(type) {
		case *models.ValidationError:
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Error de validación de negocio",
				Message: e.Error(),
				Field:   e.GetField(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error interno",
				Message: err.Error(),
			})
		}
		return
	}

	response := dto.ToProductResponseFromModel(product)
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Message: "Producto creado exitosamente",
		Data:    response,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	product, err := h.productService.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "Producto no encontrado",
			Message: err.Error(),
		})
		return
	}

	response := dto.ToProductResponseFromModel(product)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.productService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Error al listar productos",
			Message: err.Error(),
		})
		return
	}

	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ToProductResponseFromModel(&p))
	}

	c.JSON(http.StatusOK, dto.ProductListResponse{
		Products: productResponses,
		Total:    len(productResponses),
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Error de validación",
			Message: err.Error(),
		})
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), &req)
	if err != nil {
		switch e := err.(type) {
		case *models.ValidationError:
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Error:   "Error de validación de negocio",
				Message: e.Error(),
				Field:   e.GetField(),
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error:   "Error interno",
				Message: err.Error(),
			})
		}
		return
	}

	response := dto.ToProductResponseFromModel(product)
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Producto actualizado exitosamente",
		Data:    response,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "ID inválido",
			Message: "El ID debe ser un número entero positivo",
		})
		return
	}

	err = h.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "Producto no encontrado",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Producto eliminado exitosamente",
	})
}

func (h *ProductHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/products", h.CreateProduct)
	router.GET("/products", h.ListProducts)
	router.GET("/products/:id", h.GetProduct)
	router.PUT("/products/:id", h.UpdateProduct)
	router.DELETE("/products/:id", h.DeleteProduct)
}

// === ARCHIVO: internal/services/product_service.go ===
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

// === ARCHIVO: internal/utils/validator.go ===
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


// === ARCHIVO: internal/handlers/error_handler.go ===
package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/empresa/catalogo-productos/internal/dto"
)

type ErrorHandler struct {
	logger *log.Logger
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		logger: log.Default(),
	}
}

func (h *ErrorHandler) HandleValidationError(c *gin.Context, err error) {
	errMsg := err.Error()
	var fieldName string

	if strings.Contains(errMsg, "Name") {
		fieldName = "name"
	} else if strings.Contains(errMsg, "Price") {
		fieldName = "price"
	} else if strings.Contains(errMsg, "Stock") {
		fieldName = "stock"
	} else if strings.Contains(errMsg, "Category") {
		fieldName = "category"
	} else {
		fieldName = "field"
	}

	c.JSON(http.StatusBadRequest, dto.ErrorResponse{
		Success: false,
		Message: "Error de validación",
		Error:   errMsg,
		Field:   fieldName,
	})
}

func (h *ErrorHandler) HandleNotFound(c *gin.Context, resource string) {
	c.JSON(http.StatusNotFound, dto.ErrorResponse{
		Success: false,
		Message: resource + " no encontrado",
		Error:   "El recurso solicitado no existe en el sistema",
	})
}

func (h *ErrorHandler) HandleConflict(c *gin.Context, field string, value string) {
	c.JSON(http.StatusConflict, dto.ErrorResponse{
		Success: false,
		Message: "Conflicto de datos",
		Error:   "El valor '" + value + "' ya existe para el campo '" + field + "'",
		Field:   field,
	})
}

func (h *ErrorHandler) HandleBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, dto.ErrorResponse{
		Success: false,
		Message: "Solicitud incorrecta",
		Error:   message,
	})
}

func (h *ErrorHandler) HandleInternalError(c *gin.Context, err error) {
	h.logger.Printf("Error interno: %v", err)
	c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
		Success: false,
		Message: "Error interno del servidor",
		Error:   "Ocurrió un error inesperado. Por favor, contacte al administrador.",
	})
}

func (h *ErrorHandler) RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				h.logger.Printf("Panic recuperado en %s: %v", c.Request.URL.Path, err)
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
					Success: false,
					Message: "Error crítico del servidor",
					Error:   "Se produjo un error inesperado",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

// HandleMethodNotAllowed responde cuando se usa un método HTTP no permitido
func (h *ErrorHandler) HandleMethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, dto.ErrorResponse{
		Success: false,
		Message: "Método no permitido",
		Error:   "El método HTTP " + c.Request.Method + " no es válido para este endpoint",
	})
}

// HandleInvalidJSON responde cuando el JSON del body es inválido
func (h *ErrorHandler) HandleInvalidJSON(c *gin.Context) {
	c.JSON(http.StatusBadRequest, dto.ErrorResponse{
		Success: false,
		Message: "JSON inválido",
		Error:   "El cuerpo de la solicitud no es un JSON válido",
	})
}

// === ARCHIVO: config/database.go ===
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     parseIntEnv("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "catalogo_productos"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func GetDSN() string {
	cfg := getDatabaseConfig()
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func InitDatabase() (*gorm.DB, error) {
	dsn := GetDSN()

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener conexión de la base de datos: %w", err)
	}

	maxOpenConns := parseIntEnv("DB_MAX_OPEN_CONNS", 25)
	maxIdleConns := parseIntEnv("DB_MAX_IDLE_CONNS", 10)
	connMaxLifetime := parseIntEnv("DB_CONN_MAX_LIFETIME", 300)

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	log.Println("Conexión a PostgreSQL establecida correctamente")
	return db, nil
}

func getConnectionString() string {
	return GetDSN()
}

func ValidateConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	return nil
}

// === ARCHIVO: migrations/000001_init_products.up.sql ===
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


// === ARCHIVO: internal/dto/product_dto.go ===
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

// === ARCHIVO: config/database.go ===
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     parseIntEnv("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "catalogo_productos"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func GetDSN() string {
	cfg := getDatabaseConfig()
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func InitDatabase() (*gorm.DB, error) {
	dsn := GetDSN()

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error al obtener conexión de la base de datos: %w", err)
	}

	maxOpenConns := parseIntEnv("DB_MAX_OPEN_CONNS", 25)
	maxIdleConns := parseIntEnv("DB_MAX_IDLE_CONNS", 10)
	connMaxLifetime := parseIntEnv("DB_CONN_MAX_LIFETIME", 300)

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	log.Println("Conexión a PostgreSQL establecida correctamente")
	return db, nil
}

func getConnectionString() string {
	return GetDSN()
}

func ValidateConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	return nil
}

```
