package handler

import (
	"github.com/francoispqt/onelog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"kmt1/config"
	"kmt1/model"
	"os"
)

var log *onelog.Logger

func init() {
	log = onelog.New(
		os.Stdout,
		onelog.ALL,
	)
}

// inject dependencies
type Server struct {
	App   *fiber.App
	Store *config.Stor
}

func (s *Server) Start() {
	s.Routes()
	listenAddr := os.Getenv(config.ListenAddr)
	log.Fatal(s.App.Listen(listenAddr).Error())
}

func NewServer() *Server {
	stor := config.InitStore()
	model.InitDB(stor)
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	return &Server{
		App:   app,
		Store: stor,
	}
}
