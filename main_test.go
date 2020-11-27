package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"io"
	"kmt1/config"
	"kmt1/handler"
	"kmt1/model"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var URL string

func map2json(m M.SX) string {
	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println(`ERR map2json: ` + err.Error())
	}
	return string(json)
}

func json2arr(r io.Reader) A.MSX {
	res := A.MSX{}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		fmt.Println(`ERR json2arr: ` + err.Error())
	}
	return res
}

func json2map(r io.Reader) M.SX {
	res := M.SX{}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		fmt.Println(`ERR json2map: ` + err.Error())
	}
	return res
}

func compareString(t *testing.T, key, in, out string) {
	if in != out {
		t.Fatal(`different ` + key + `: ` + in + ` <> ` + out)
	}
}

func compareAny(t *testing.T, key string, in, out interface{}) {
	i := X.ToS(in)
	o := X.ToS(out)
	if i != o {
		t.Fatal(`different ` + key + `: ` + i + ` <> ` + o)
	}
}

func postApi(t *testing.T, path string, jsonIn interface{}, c *http.Client, headers M.SS) M.SX {
	fmt.Printf("Hitting API: %s %v\n", path, jsonIn)
	body := strings.NewReader(X.ToJson(jsonIn))
	req, err := http.NewRequest(`POST`, URL+path, body)
	if err != nil {
		t.Error(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}
	jsonOut := json2map(res.Body)
	fmt.Printf("%#v\n", jsonOut)
	return jsonOut
}

func getApi(t *testing.T, path string, c *http.Client, param M.SS, headers M.SS) A.MSX {
	req, err := http.NewRequest(`GET`, URL+path+param2querystring(param), nil)
	if err != nil {
		t.Error(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := c.Do(req)
	if err != nil {
		t.Error(err)
	}
	jsonOut := json2arr(res.Body)
	fmt.Printf("%#v\n", jsonOut)
	return jsonOut
}

func param2querystring(param M.SS) string {
	res := ``
	for k, v := range param {
		if res == `` {
			res += `?`
		} else {
			res += `&`
		}
		res += k + `=` + S.Replace(v, ` `, `%20`)
	}
	return res
}

func TestApis(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	URL = `http://` + os.Getenv(config.ListenAddr)
	config.LoadEnv()
	c := &http.Client{}

	// HIT ARTICLE create
	articleIn := model.Article{
		Title:  faker.Sentence(),
		Author: faker.Name(),
		Body:   faker.Paragraph(),
	}
	articleOut := postApi(t, handler.Article, articleIn, c, nil)
	L.Describe(articleOut)

	articleId := articleOut.GetInt(`id`)
	if articleId <= 0 {
		t.Fatal(`failed create`)
	}

	L.Describe(articleOut)
	compareString(t, `article title`, articleIn.Title, articleOut.GetStr(`title`))
	compareString(t, `article author`, articleIn.Author, articleOut.GetStr(`author`))
	compareString(t, `article body`, articleIn.Body, articleOut.GetStr(`body`))

	// HIT ARTICLE search
	searchIn := M.SS{
		`query`: `aaaaaaaaaaaaaaaaaaaaaaaaa`,
	}

	searchOut := getApi(t, handler.Article, c, searchIn, nil)
	L.Describe(searchOut)
}
