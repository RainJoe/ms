package config

import (
        "github.com/BurntSushi/toml"
        "os"
)

var (
        //Cfg is user config
        Cfg     = &Config{Name: "gateway"}
        cfgPath = "./config/gateway-dev.toml"
)

//Config def
type Config struct {
        Name       string
        Port       string
        Version    string
        UserSeviceAddr userServiceConfig `toml:"userService"`
}

type userServiceConfig struct {
        Addr string
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
                _, err = toml.Decode(os.Getenv("GATEWAY_CONFIG"), &Cfg)
        } else {
                _, err = toml.DecodeFile(cfgPath, &Cfg)
        }
        return
}
