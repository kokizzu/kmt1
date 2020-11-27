package model

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"
	"kmt1/model/adapter"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const CacheSpace = `caches`

const Key = `key`
const Val = `val`

var CacheTable = &adapter.TableProp{
	Droppable: true,
	Fields: []adapter.Field{
		{Name: Key, Type: adapter.String},
		{Name: Val, Type: adapter.String},
	},
	Unique: Key,
}

type Cache struct {
	Key string
	Val string
}

func (c *Cache) Set(s *adapter.Taran, cacheKey string, cacheVal string) {
	res, err := s.Upsert(CacheSpace, A.X{cacheKey, cacheVal}, A.X{})
	L.IsError(err, `failed to set cache`)
	L.Describe(res.Tuples())
}

func (c *Cache) ClearAll(s *adapter.Taran) {
	s.CallBoxSpace(CacheSpace+`:truncate`, A.X{})
}

func (c *Cache) Get(s *adapter.Taran, cacheKey string) string {
	res, err := s.Select(CacheSpace, Key, 0, 1, tarantool.IterEq, []interface{}{cacheKey})
	if L.IsError(err, `failed to get from cache`) {
		return `[]`
	}
	matrix := res.Tuples()
	if len(matrix) == 0 {
		return `[]`
	}
	row := matrix[0]
	if len(row) == 0 {
		return `[]`
	}
	return X.ToS(row[0])
}
