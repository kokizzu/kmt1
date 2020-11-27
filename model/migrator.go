package model

import (
	"kmt1/config"
)

func InitDB(s *config.Stor) error {

	// tarantool
	s.Taran.MigrateTarantool(ArticleSpace, ArticleTable)
	s.Taran.MigrateTarantool(CacheSpace, CacheTable)
	c := Cache{}
	// clear all cache
	c.ClearAll(s.Taran)

	// meilisearch
	rankingRules := []string{
		`desc(created)`,
		`typo`,
		`words`,
		`proximity`,
		`attribute`,
		`wordsPosition`,
		`exactness`,
	}
	if err := s.Meili.MigrateMeilisearch(ArticleSpace, Id, rankingRules); err != nil {
		return err
	}
	return nil
}
