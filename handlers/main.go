package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/GreenHedgehog/solar-test/db"
	"github.com/GreenHedgehog/solar-test/models"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func add(c echo.Context) error {

	vacancy := new(models.Vacancy)

	if err := c.Bind(vacancy); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := db.Add(vacancy); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, vacancy)
}

func delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := db.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	v, err := db.Get(id)

	switch err {
	case sql.ErrNoRows:
		return c.NoContent(http.StatusNotFound)
	case nil:
		return c.JSON(http.StatusOK, v)
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func getAll(c echo.Context) error {
	data, err := db.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, data)
}

func login(c echo.Context) error {

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if user.Login == "login" && user.PassHash == "hash" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Add hooks up all endpoints
func Add(e *echo.Echo) {

	e.POST("/login", login)

	e.PUT("/vacancy", add, middleware.JWT([]byte("secret")))
	e.DELETE("/vacancy/:id", delete, middleware.JWT([]byte("secret")))
	e.GET("/vacancy/:id", get)
	e.GET("/vacancy", getAll)
}
