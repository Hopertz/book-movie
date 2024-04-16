package main

import (
	"errors"
	"log/slog"
	"movie-booking/pkg/mongodb"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Amount int32              `json:"amount,omitempty" bson:"amount,omitempty"`
	Seats  [][]string         `json:"seats,omitempty" bson:"seats,omitempty"`
}

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

	var res []*mongodb.Movie

	ctx := c.Request().Context()

	cursor, err := app.models.Movie.Find(ctx, bson.M{}, options.Find().SetProjection(bson.M{"id": 1, "name": 1, "amount": 1}))

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

		res = append(res, &elem)
	}

	return c.JSON(http.StatusOK, res)

}

func (app *application) getMovieById(c echo.Context) error {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var movie mongodb.Movie

	ctx := c.Request().Context()

	err = app.models.Movie.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movie)

}
