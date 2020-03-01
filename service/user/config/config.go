package config

import (
	"os"
	"github.com/BurntSushi/toml"
	"github.com/RainJoe/ms/library/pg"
	"github.com/RainJoe/ms/library/redis"
)

var (
	//Cfg is user config
	Cfg     = &Config{Name: "ms"}
	cfgPath = "./config/ms-dev.toml"
)

//Config def
type Config struct {
	Name  string
	Pg    pg.Config   `toml:"pg"`
	Redis redisConfig `toml:"redis"`
	Jwt   jwtConfig   `toml:"jwt"`
}

type jwtConfig struct {
	Key string
}

type redisConfig struct {
	redis.Config
	TokenKey string
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//Init config
func Init() (err error) {
	if !exists(cfgPath) {
		_, err = toml.Decode(os.Getenv("MS_CONFIG"), &Cfg)
	} else {
		_, err = toml.DecodeFile(cfgPath, &Cfg)
	}
	return
}
