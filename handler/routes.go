package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
)

const (
	DEBUG   = true
	Article = `/articles`
)

// inject server
type Ctx struct {
	*Server
	*fiber.Ctx
}

func (s *Server) Handler(handler func(ctx *Ctx) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if DEBUG {
			debugPrefix := string(c.Context().Method()) + ` ` + c.OriginalURL() + ` `
			L.Print(debugPrefix + string(c.Body()))
			defer (func() {
				L.Print(debugPrefix + string(c.Response().Body()))
			})()
		}
		return handler(&Ctx{
			s,
			c,
		})
	}
}

func (s *Server) Routes() {

	s.App.Post(Article, s.Handler(PostArticle))
	s.App.Get(Article, s.Handler(GetArticle))

}
