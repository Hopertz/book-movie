package main

import (
	"context"
	"log/slog"
	"movie-booking/pkg/mongodb"
	"net/http"
	"strconv"
	"time"

	"github.com/Golang-Tanzania/mpesa"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (app *application) insertBooker(c echo.Context) error {

	var booker struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Seat  []struct {
			Row  int `json:"row"`
			Seat int `json:"seat"`
		} `json:"seat"`

		Movie  string     `json:"movie"`
		Amount int32      `json:"amount"`
		Seats  [][]string `json:"seats"`
	}

	if err := c.Bind(&booker); err != nil {
		return err
	}

	s := mongodb.SelectedSeat{
		Row:  booker.Seat[0].Row,
		Seat: booker.Seat[0].Seat,
	}

	// loop through booker.Seats and set from selected to ocupied
	for i := 0; i < len(booker.Seats); i++ {
		for j := 0; j < len(booker.Seats[i]); j++ {
			if booker.Seats[i][j] == "selected" {
				booker.Seats[i][j] = "occupied"
			}
		}
	}

	now := time.Now()

	// Convert current date and time to primitive.DateTime
	dateTime := primitive.DateTime(now.UnixNano() / int64(time.Millisecond))

	b := mongodb.Booker{
		Name:      booker.Name,
		Phone:     booker.Phone,
		Seat:      []mongodb.SelectedSeat{s},
		CreatedAt: primitive.DateTime(dateTime),
	}

	m := mongodb.Movie{
		Name:   booker.Movie,
		Amount: booker.Amount,
		Seats:  booker.Seats,
	}

	txid, _ := app.generateRandomString()

	price := strconv.Itoa(len(booker.Seats) * int(booker.Amount))

	phone := booker.Phone

	no := "255" + phone[1:]

	payload := mpesa.C2BPaymentRequest{
		Amount:                   price,
		CustomerMSISDN:           no,
		Country:                  "TZN",
		Currency:                 "TZS",
		ServiceProviderCode:      "000000",
		TransactionReference:     "T12344C",
		ThirdPartyConversationID: txid,
		PurchasedItemsDesc:       "Test",
	}

	_, err := app.mpesaClient.C2BPayment(context.Background(), payload)

	if err != nil {
		slog.Error("err", "errorr in making mpesa txn ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// update movie seats
	_, err = app.models.Movie.UpdateOne(c.Request().Context(), bson.M{"name": m.Name}, bson.M{"$set": m})
	if err != nil {
		slog.Error("err", "errorr in updating movie", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// insert booker
	_, err = app.models.Booker.InsertOne(c.Request().Context(), b)

	if err != nil {
		slog.Error("err", "error inserting booker", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, booker)

}
