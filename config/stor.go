package config

import (
	"github.com/francoispqt/onelog"
	"github.com/meilisearch/meilisearch-go"
	"github.com/tarantool/go-tarantool"
	"kmt1/model/adapter"
	"os"
)

type Stor struct {
	Taran *adapter.Taran
	Meili *adapter.Meili
	Log   *onelog.Logger
}

func InitStore() *Stor {
	s := newStore()
	taranUser := os.Getenv(TaranUser)
	opts := tarantool.Opts{User: taranUser}
	taranHost := os.Getenv(TaranHost)
	taranConn, err := tarantool.Connect(taranHost, opts)
	if err != nil {
		log.Fatal(`cannot connect to tarantool: ` + taranUser + `@` + taranHost)
	}
	meiliHost := os.Getenv(MeiliHost)
	meiliKey := os.Getenv(MeiliKey)
	meiliClient := meilisearch.NewClient(meilisearch.Config{
		Host:   meiliHost,
		APIKey: meiliKey,
	})

	s.Taran = &adapter.Taran{taranConn, s.Log}
	s.Meili = &adapter.Meili{meiliClient, s.Log}
	return s
}

func newStore() *Stor {
	return &Stor{
		Log: log,
	}
}
