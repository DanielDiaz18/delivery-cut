package handlers

import (
	"net/http"

	"github.com/DanielDiaz18/rappi-cut/backend/data"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (s *ServerHandler) CreateParner(c echo.Context) error {
	form := &data.ParnerForm{}

	if c.Bind(form) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid form")
	}

	parner := &data.Parner{
		UserId: form.UserId,
	}

	if err := s.p.Create(parner); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, parner)
}

func (s *ServerHandler) GetParners(c echo.Context) error {
	kind := c.Get("user").(*jwt.Token).Claims.(*data.CustomClaims).Kind
	if kind != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
	}
	prs := data.Parners{}
	if err := s.p.GetAll(prs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, prs)
}
