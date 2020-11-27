package model

import (
	"kmt1/config"
)

func InitDB(s *config.Stor) error {

	// tarantool
	s.Taran.MigrateTarantool(ArticleSpace, ArticleTable)
	s.Taran.MigrateTarantool(CacheSpace, CacheTable)

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
