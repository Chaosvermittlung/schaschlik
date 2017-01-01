package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"fmt"

	"github.com/chaosvermittlung/schaschlik/config"
	"github.com/chaosvermittlung/schaschlik/irc"
	"github.com/chaosvermittlung/schaschlik/writer"
)

var configpath = flag.String("configpath", "", "Path to the config (default is executable folder)")

func main() {
	flag.Parse()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	execdir := dir + "/"
	if *configpath == "" {
		*configpath = execdir + "config.json"
	}
	fmt.Println(*configpath)
	conf, err := config.LoadConfig(*configpath)
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	writer.Setup(conf)
	//server.Setup(conf)
	irc.Setup(conf.IRCs)
	for {
	}
}
