package internal

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olukkas/ushort/internal/controllers"
	"log"
	"os"
)

type App struct {
	server   *fiber.App
	database *sql.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {
	a.setupDb()
	a.setupRoutes()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	log.Fatal(a.server.Listen(":" + port))
}

func (a *App) Close() error {
	if a.database != nil {
		return a.database.Close()
	}

	return nil
}

func (a *App) setupRoutes() {
	a.server = fiber.New(fiber.Config{})
	a.server.Use(cors.New())
	a.server.Use(logger.New())

	a.server.Get("/", controllers.HelloRoute)
	a.server.Static("/static", "./static")
}

func (a *App) setupDb() {
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("DB_NAME environment variable not set")
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	a.database = db
}
