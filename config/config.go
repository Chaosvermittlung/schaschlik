package config

import (
	"encoding/json"
	"io/ioutil"
)

//IRCConn is the struct for IRC Connection data
type IRCConn struct {
	Nick       string
	Username   string
	Adress     string
	Channel    string
	TLS        bool
	SkipVerify bool
}

//Config is the struct to configure the Server
type Config struct {
	Type string
	Port string
	Path string
	Size int
	IRCs []IRCConn
}

//LoadConfig loads a Config from a Path and gives back the Config struct
func LoadConfig(path string) (Config, error) {
	var conf Config
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(body, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
