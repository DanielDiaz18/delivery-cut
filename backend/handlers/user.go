package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DanielDiaz18/rappi-cut/backend/data"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (u *ServerHandler) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*data.CustomClaims)
	us, err := u.r.Get(claims.Id.Hex())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, us)
}

func (h *ServerHandler) Login(c echo.Context) error {
	form := &data.LoginForm{}

	if c.Bind(form) != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid form")
	}

	u := &data.User{}
	if err := h.r.Login(form.Email, form.Password, u); err != nil {
		return err
	}

	jwtClaims := &data.CustomClaims{
		Id:   u.ID,
		Kind: u.Kind,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return err
	}

	tkn, err := jwt.ParseWithClaims(t, &data.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)

	}
	log.Println(tkn.Claims.(*data.CustomClaims))

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user":  u,
	})
}

func (h *ServerHandler) CreateUser(c echo.Context) error {
	u := &data.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.HashPassword(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := h.r.Create(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if u.Kind != "user" {
		token := c.Get("user")
		if token == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		user := token.(*jwt.Token).Claims.(*data.CustomClaims)
		if user == nil || user.Kind != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		if user.Kind == "parner" {
			p := &data.Parner{
				UserId: user.Id,
			}
			if err := h.p.Create(p); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
	}
	return c.JSON(http.StatusCreated, u)
}

func (u *ServerHandler) HashPassword(user *data.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
