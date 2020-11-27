package config

import (
	"github.com/francoispqt/onelog"
	"github.com/joho/godotenv"
	"github.com/kokizzu/gotro/L"
	"os"
)

const ConfigFile = `config.env`
const TaranHost = `TARAN_HOST`
const TaranUser = `TARAN_USER`
const ListenAddr = `LISTEN_ADDR`
const MeiliHost = `MEILI_HOST`
const MeiliKey = `MEILI_KEY`

var log *onelog.Logger

func init() {
	log = onelog.New(
		os.Stdout,
		onelog.ALL,
	)
}

func LoadEnv() {
	log.Info(`loading ` + ConfigFile)
	err := godotenv.Load(ConfigFile)
	if L.IsError(err, `failed to load `+ConfigFile) {
		log.Error(err.Error())
		return
	}
}
