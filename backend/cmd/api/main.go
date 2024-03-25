package main

import (
	"movie-booking/pkg/mongodb"
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"sync"
	"time"
)

type config struct {
	port int
	env  string

	db struct {
		dsn string
	}
}

type application struct {
	config config
	wg     sync.WaitGroup
	models *mongodb.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8448, "API server port")
	flag.StringVar(&cfg.env, "env", "production", "Environment (development|Staging|production")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "MongoDB DSN")

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

	app := &application{
		config: cfg,
		models: mongodb.NewModels(cli),
	}

	slog.Info("Starting server on", "port", cfg.port)
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}
