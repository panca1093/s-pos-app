package product

import (
	"github.com/s-pos-app/internal/utilities/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	usecase IUsecase
}

func NewHandler(uc IUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: uc}
}

func (h *HTTPHandler) Mount(e *echo.Echo) {
	e.GET("/products", h.GetAll)
	e.GET("/products/:id", h.GetByID)
	e.POST("/products", h.Create)
}

func (h *HTTPHandler) GetAll(c echo.Context) error {
	products, err := h.usecase.GetAll(c.Request().Context())
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to fetch products")
	}

	return response.Success(c, http.StatusOK, products)
}

func (h *HTTPHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return response.Error(c, http.StatusBadRequest, "invalid product ID")
	}
	// Convert id to uint
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid product ID")
	}

	products, err := h.usecase.GetByID(c.Request().Context(), productID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to fetch products")
	}

	return response.Success(c, http.StatusOK, products)
}

func (h *HTTPHandler) Create(c echo.Context) error {
	var p IProduct
	if err := c.Bind(&p); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid input")
	}

	if err := c.Validate(&p); err != nil {
		return response.Error(c, http.StatusBadRequest, "validation failed")
	}

	// Assuming the database auto-generates the product ID
	product, err := h.usecase.Create(c.Request().Context(), p)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "failed to create product")
	}

	return response.Success(c, http.StatusCreated, product)
}
