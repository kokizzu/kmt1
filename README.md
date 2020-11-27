
# KMT1

evaluating meilisearch and tarantool with go

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
- [x] add DEBUG flag to print debugging handler
- [ ] add DEBUG flag for model
- [ ] add autorecompile
- [ ] add autoupdater
- [ ] add requeue when meilisearch failed
- [ ] add bucketed cache (so no need to expire all cache when record upserted)
- [ ] bugfix error migration if running first time [github-issue](https://github.com/tarantool/go-tarantool/issues/94)
- [ ] make functions to simplify CRUD operations 
- [ ] add expiration daemon (so can have TTL like redis/aerospike)
- [ ] expire the search cache
- [ ] wait for index to be created in test if test executed for the first time [github-issue](https://github.com/meilisearch/meilisearch-go/issues/108), workaround: create an empty document for the first time (so search engine know what are the columns)
- [ ] try [sonic](https://github.com/valeriansaliou/sonic) search engine

## dependencies

- [fiber](https://gofiber.io/) framework that use one of the fastest router: fasthttp)
- [tarantool](https://www.tarantool.io/en/) works as cache and also DBMS (support SQL and WAL by default) automatically caches per record -- not used in this case, because the read/query part fetched from search engine
- [meilisearch]() (one of the fastest search engine for small data sets)
- [jsoniter](https://github.com/json-iterator/go) (fastest std-compatible json converter)
