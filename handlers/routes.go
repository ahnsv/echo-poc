package handlers

import "github.com/labstack/echo"

func (h *Handler) Register(v1 *echo.Group) {
	users := v1.Group("/users")
	users.GET("", h.getUser)
	users.POST("", h.createUser)
}
