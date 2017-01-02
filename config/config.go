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

//Server is the struct to define a tcp or udp Listener to send messages to
type Server struct {
	Type string
	Port string
	Size int
}

//Config is the struct to configure the Server
type Config struct {
	Path    string
	Servers []Server
	IRCs    []IRCConn
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
