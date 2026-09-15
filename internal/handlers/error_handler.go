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