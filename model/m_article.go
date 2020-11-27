package model

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/meilisearch/meilisearch-go"
	"kmt1/config"
	"kmt1/model/adapter"
	"time"
)

const ArticleSpace = `articles`
const Id = `id`
const Author = `author`
const Title = `title`
const Body = `body`
const Created = `created`

var ArticleTable = &adapter.TableProp{
	Droppable: true,
	Fields: []adapter.Field{
		{Name: Id, Type: adapter.Unsigned},
		{Name: Author, Type: adapter.String},
		{Name: Title, Type: adapter.String},
		{Name: Body, Type: adapter.String},
		{Name: Created, Type: adapter.Unsigned},
	},
	Unique:        Id,
	AutoIncrement: true,
}

type Article struct {
	Id      int64  `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Created int64  `json:"created"` // epoch
}

func (a *Article) ToMsgPack() []interface{} {
	res := []interface{}{
		a.Id,
		a.Author,
		a.Title,
		a.Body,
		a.Created,
	}
	if a.Id == 0 {
		res[0] = nil
	}
	return res
}

func (a *Article) Create(s *config.Stor) error {
	a.Id = 0
	if a.Created == 0 {
		a.Created = time.Now().UnixNano()
	}
	row := a.ToMsgPack()
	res, err := s.Taran.Insert(ArticleSpace, row)
	if L.IsError(err, `failed insert tarantool`) {
		return err
	}

	// get the id
	tup := res.Tuples()
	a.Id = X.ToI(tup[0][0])

	// clean cache
	c := Cache{}
	c.ClearAll(s.Taran)

	// index on search engine
	ret, err := s.Meili.UpsertOne(ArticleSpace, a, Id)
	L.Describe(X.ToJson(a))
	L.Describe(ret)
	L.Describe(a)
	if L.IsError(err, `failed insert meilisearch`) {
		return err
	}
	return nil
}

type ArticleSearchIn struct {
	Query  string `query:"query"`
	Author string `author:"author"`
}

func (as *ArticleSearchIn) ToMeiliQuery() (*meilisearch.SearchRequest, string) {
	authorQuery := ``
	if as.Author != `` {
		authorQuery = `author=` + S.ZZ(as.Author)
	}
	searchReq := &meilisearch.SearchRequest{
		Query:   as.Query,
		Filters: authorQuery,
	}
	cacheKey, err := json.Marshal(searchReq)
	L.IsError(err, `failed convert SearchQuery to json`)
	return searchReq, string(cacheKey)
}

func (a *Article) Search(s *config.Stor, in *ArticleSearchIn) string {
	c := Cache{}
	searchReq, cacheKey := in.ToMeiliQuery()
	if res := c.Get(s.Taran, cacheKey); res != `[]` {
		L.Print(`============== from cache`)
		return res
	}
	res, err := s.Meili.Query(ArticleSpace, searchReq)
	if L.IsError(err, `failed to search: `+cacheKey) {
		return `{"error":"failed search"}`
	}
	json, err := res.MarshalJSON()
	if L.IsError(err, `failed to convert json from search engine`) {
		return `{"error":"failed convert json"}`
	}
	if len(res.Hits) > 0 {
		c.Set(s.Taran, cacheKey, string(json))
	}
	return string(json)
}
