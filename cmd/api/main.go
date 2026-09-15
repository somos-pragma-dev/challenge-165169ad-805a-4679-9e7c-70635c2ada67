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