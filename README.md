
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

# test server (start the server first)
go test
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
- [ ] find meilisearch
- [ ] cache meilisearch
- [ ] add apitest (faker) create article
- [ ] add apitest search article negative
- [ ] add apitest search article positive
- [ ] add apitest search article cached
- [ ] add autorecompile
- [ ] add autoupdater
- [ ] add requeue when meilisearch failed

## dependencies

- fiber (one of the fastest router)
- tarantool (works as cache and also dbms)
- meilisearch (as search engine)
