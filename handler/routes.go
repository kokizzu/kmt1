package handler

import "github.com/gofiber/fiber"

const (
	Article = `/articles`
)

// inject server
type Ctx struct {
	*Server
	*fiber.Ctx
}

func (s *Server) Handler(handler func(ctx *Ctx) error) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		handler(&Ctx{
			s,
			c,
		})
	}
}

func (s *Server) Routes() {

	s.App.Post(Article, s.Handler(PostArticle))
	s.App.Get(Article, s.Handler(GetArticle))

}
