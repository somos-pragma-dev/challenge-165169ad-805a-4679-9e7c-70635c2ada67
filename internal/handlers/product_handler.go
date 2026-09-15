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