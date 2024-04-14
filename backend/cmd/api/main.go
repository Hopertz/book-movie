package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"movie-booking/pkg/mongodb"
	"os"
	"sync"
	"time"

	"github.com/Golang-Tanzania/mpesa"
)

type config struct {
	port int
	env  string

	db struct {
		dsn string
	}

	mpesa string
}

type application struct {
	config      config
	wg          sync.WaitGroup
	models      *mongodb.Models
	mpesaClient *mpesa.Client
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8448, "API server port")
	flag.StringVar(&cfg.env, "env", "production", "Environment (development|Staging|production")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "MongoDB DSN")
	flag.StringVar(&cfg.mpesa, "mpesa-api-key", os.Getenv("MPESA_API_KEY"), "MPESA_API_KEY")

	flag.Parse()

	cli, err := openDB(cfg)
	if err != nil {
		log.Fatal(err, nil)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = cli.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	slog.Info("Connected to MongoDB")

	// initialize mpesa
	// clientMpesa, err := mpesa.NewClient(cfg.mpesa, mpesa.Sandbox, 24)

	// if err != nil {
	// 	log.Fatal(err, nil)
	// }

	app := &application{
		config:      cfg,
		models:      mongodb.NewModels(cli),
		// mpesaClient: clientMpesa,
	}

	slog.Info("Starting server on", "port", cfg.port)
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}
