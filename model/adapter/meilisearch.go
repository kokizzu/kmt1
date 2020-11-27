package adapter

import (
	"github.com/francoispqt/onelog"
	"github.com/meilisearch/meilisearch-go"
)

type Meili struct {
	meilisearch.ClientInterface
	Log *onelog.Logger
}
