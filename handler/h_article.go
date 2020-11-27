package handler

import (
	"github.com/gofiber/fiber/v2"
	"kmt1/model"
)

func GetArticle(ctx *Ctx) error {
	search := model.ArticleSearchIn{}
	if err := ctx.QueryParser(&search); err != nil {
		return err
	}
	article := model.Article{}
	res := article.Search(ctx.Store, &search)
	if res == `` {
		res = `{"error":"empty result"}`
	}
	ctx.Response().Header.SetContentType(fiber.MIMEApplicationJSON)
	ctx.SendString(res)
	return nil
}

func PostArticle(ctx *Ctx) error {
	article := model.Article{}
	if err := ctx.BodyParser(&article); err != nil {
		return err
	}
	err := article.Create(ctx.Store)
	if err == nil {
		ctx.JSON(article)
	}
	return nil
}
