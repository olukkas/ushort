package main

import (
	"github.com/joho/godotenv"
	"github.com/olukkas/ushort/internal"
	"log/slog"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found, using system defaults")
	}

	app := internal.NewApp()
	app.Init()

	err := app.Close()
	if err != nil {
		panic(err.Error())
	}
}
