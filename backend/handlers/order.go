package handlers

import (
	"net/http"

	"github.com/DanielDiaz18/rappi-cut/backend/data"
	"github.com/golang-jwt/jwt"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

func (s *ServerHandler) CreateOrder(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	if token == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
	}
	claims := token.Claims.(*data.CustomClaims)

	form := &data.OrderForm{}

	if c.Bind(form) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid form")
	}

	rou := &data.Route{
		Origin:      form.Origin,
		Destination: form.Destination,
	}

	if err := mgm.Coll(rou).Create(rou); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid form")
	}

	ord := &data.Order{
		Route:  rou.ID,
		User:   claims.Id,
		Parner: form.Parner,
		Status: "initial",
	}
	if err := s.o.Create(ord); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ord)
}

func (s *ServerHandler) GetOrders(c echo.Context) error {
	user := c.Get("user")
	if user == nil || user.(*jwt.Token).Claims.(*data.CustomClaims).Kind != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "You are not authorized")
	}
	ords := data.Orders{}
	if err := s.o.GetOrders(ords); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ords)
}
