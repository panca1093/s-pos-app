package user

import (
	"github.com/labstack/echo/v4"
	"github.com/s-pos-app/internal/utilities/response"
	"net/http"
)

type HTTPHandler struct {
	usecase IUsecase
}

func NewHandler(uc IUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: uc}
}

func (h *HTTPHandler) Mount(e *echo.Echo) {
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}

func (h *HTTPHandler) Register(c echo.Context) error {
	var u IUser
	if err := c.Bind(&u); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid input")
	}

	// Validate the input
	if u.Username == "" || u.Password == "" {
		return response.Error(c, http.StatusBadRequest, "username and password are required")
	}

	if err := h.usecase.Register(&u); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusCreated, "user created")
}

func (h *HTTPHandler) Login(c echo.Context) error {
	var in struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&in); err != nil {
		return response.Error(c, http.StatusBadRequest, "invalid input")
	}

	usr, token, err := h.usecase.Login(in.Username, in.Password)
	if err != nil {
		return response.Error(c, http.StatusUnauthorized, err.Error())
	}

	if usr == nil {
		return response.Error(c, http.StatusUnauthorized, "user not found")
	}

	data := map[string]interface{}{
		"id":       usr.ID,
		"name":     usr.Name,
		"username": usr.Username,
		"role":     usr.Role,
		"token":    token,
	}
	return response.Success(c, http.StatusOK, data)
}
