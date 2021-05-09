package handlers

import (
	"echo-poc/models"
	"net/http"

	"github.com/labstack/echo"
)

type userCreationRequest struct {
	Username string `json:"username" validate:"required"`
}

type userResponse struct {
	Username string `json:"username" validate:"required"`
}

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.Username = u.Name
	return r
}

func (r *userCreationRequest) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	u.Name = r.Username
	return nil
}

func (h *Handler) getUser(c echo.Context) error {
	c.Logger().Info("hellooooooo")
	return c.String(http.StatusOK, "hello")
}

func (h *Handler) createUser(c echo.Context) error {
	var u models.User
	req := &userCreationRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusCreated, newUserResponse(&u))
}
