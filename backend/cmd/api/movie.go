package main

import (
	"errors"
	"log/slog"
	"movie-booking/pkg/mongodb"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

func (app *application) insertMovie(c echo.Context) error {

	var Movie struct {
		Name   string     `json:"name" `
		Amount int32      `json:"amount"`
		Seats  [][]string `json:"seats"`
	}

	if err := c.Bind(&Movie); err != nil {
		return err
	}

	m := mongodb.Movie{
		Name:   Movie.Name,
		Amount: Movie.Amount,
		Seats:  Movie.Seats,
	}

	// check if name already exists
	var result mongodb.Movie

	err := app.models.Movie.FindOne(c.Request().Context(), bson.M{"name": m.Name}).Decode(&result)

	if err == nil {
		return c.JSON(http.StatusBadRequest, errors.New("movie already exist"))
	} else {
		// insert movie
		_, err := app.models.Movie.InsertOne(c.Request().Context(), m)
		if err != nil {
			slog.Error("err", "errorr in inserting movie", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, Movie)
}

func (app *application) getMovies(c echo.Context) error {

	var results []*mongodb.Movie

	ctx := c.Request().Context()

	cursor, err := app.models.Movie.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem mongodb.Movie
		err := cursor.Decode(&elem)
		if err != nil {
			slog.Error("err", "error getting movies", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		results = append(results, &elem)
	}

	return c.JSON(http.StatusOK, results)

}
