
# README

test meilisearch and tarantool using go

```
cd $GOPATH
git clone git@gitlab.com:kokizzu/kmt1.git kmt1 # or
git clone https://gitlab.com/kokizzu/kmt1.git kmt1

# setup infra
docker-compose up 

# run server 
go run main.go

# test APIs (start the server first)
go test

# access tarantool
tarantoolctl connect localhost:3302

# check search engine update status
curl -X GET 'http://localhost:7700/indexes/articles/updates'
```

## 

- [x] docker compose
- [x] setup router
- [x] setup store and config
- [x] connect tarantool
- [x] connect meilisearch
- [x] init schema tarantool article
- [x] init schema tarantool cache
- [x] init schema meilisearch article
- [x] insert tarantool
- [x] insert meilisearch
- [x] find using meilisearch
- [x] cache meilisearch result
- [x] unset cache when insert
- [x] add apitest (faker) create article
- [x] add apitest search article negative title/body word
- [x] add apitest search article negative author word
- [x] add apitest search article positive title
- [x] add apitest search article positive body
- [x] add apitest search article positive author
- [x] add apitest search article cached
- [ ] add autorecompile
- [ ] add autoupdater
- [ ] add requeue when meilisearch failed
- [ ] add bucketed cache (so no need to expire all cache when record upserted)
- [ ] bugfix error migration if running first time [issue-ref](https://github.com/tarantool/go-tarantool/issues/94)
- [ ] make functions to simplify CRUD operations 
- [x] set flag to print debugging (handler.DEBUG)
- [ ] add expiration daemon (so can have TTL like redis/aerospike)
- [ ] expire the search cache
- [ ] wait for index to be created in test if test executed for the first time [issue-ref](https://github.com/meilisearch/meilisearch-go/issues/108), workaround: create an empty document for the first time (so search engine know what's the columns)
- [ ] add more DEBUG flags before production

## dependencies

- fiber (framework that use one of the fastest router/fasthttp)
- tarantool (works as cache and also dbms -- automatically caches per record -- not used in this case, because using search engine)
- meilisearch (one of the fastest single server search engine)
- jsoniter (fastest std-compatible json converter)
