package main

import (
	"movie-booking/pkg/mongodb"
	"net/http"

	"github.com/labstack/echo"
)

func (app *application) createMovie(c echo.Context) error {

	var Movie struct {
		Name  string     `json:"name" `
		Amount int32      `json:"amount"`
		Seats [][]string `json:"seats"`
	}

	if err := c.Bind(&Movie); err != nil {
		return err
	}

	m := mongodb.Movie{
		Name:  Movie.Name,
		Amount: Movie.Amount,
		Seats: Movie.Seats,
	}

	if _, err := app.models.Movie.InsertOne(c.Request().Context(), m); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, Movie)
}
