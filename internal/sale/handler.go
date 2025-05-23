package sale

import "github.com/labstack/echo/v4"

type HTTPHandler struct {
	usecase IUsecase
}

func NewHandler(uc IUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: uc}
}

func (h *HTTPHandler) Mount(e *echo.Echo) {}
