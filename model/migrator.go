package model

import (
	"github.com/kokizzu/gotro/L"
	"github.com/meilisearch/meilisearch-go"
	"kmt1/config"
)

func InitDB(s *config.Stor) error {
	// tarantool
	s.Taran.MigrateTarantool(ArticleSpace, ArticleTable)
	// TODO: init tarantool schema
	// meilisearch
	_, err := s.Meili.Indexes().Create(meilisearch.CreateIndexRequest{
		UID: ArticleSpace,
	})
	L.IsError(err, `failed create index`)
	return err
}
