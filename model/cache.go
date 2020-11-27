package model

import "kmt1/model/adapter"

const CacheSpace = `caches`

const Key = `key`
const Val = `val`

type Cache struct {
	Key string
	Val string
}

var CacheTable = &adapter.TableProp{
	Droppable: true,
	Fields: []adapter.Field{
		{Name: Key, Type: adapter.String},
		{Name: Val, Type: adapter.String},
	},
	Unique: Key,
}

func (c *Cache) Set() {

}

func (c *Cache) ClearAll() {

}
