package model

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"kmt1/config"
	"kmt1/model/adapter"
)

const ArticleSpace = `articles`
const Id = `id`
const Author = `author`
const Title = `title`
const Body = `body`
const Created = `created`

type Article struct {
	Id      int64
	Author  string
	Title   string
	Body    string
	Created int64
}

var ArticleTable = &adapter.TableProp{
	Droppable: true,
	Fields: []adapter.Field{
		{Name: Id, Type: adapter.Unsigned},
		{Name: Author, Type: adapter.Unsigned},
		{Name: Title, Type: adapter.String},
		{Name: Body, Type: adapter.String},
		{Name: Created, Type: adapter.Unsigned},
	},
	Unique: Id,
}

func (a *Article) Create(s *config.Stor) error {
	a.Id = 0
	res, err := s.Taran.Insert(ArticleSpace, *a)
	if L.IsError(err, `failed insert tarantool`) {
		return err
	}
	tup := res.Tuples()
	L.Describe(tup)
	_, err = s.Meili.Documents(ArticleSpace).AddOrReplaceWithPrimaryKey(a, X.ToS(a.Id))
	if L.IsError(err, `failed insert meilisearch`) {
		return err
	}
	return nil
}

type ArticleSearchIn struct {
	Query  string
	Author string
}

func (a *Article) Search(s *config.Stor) {

}
