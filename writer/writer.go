package writer

import (
	"fmt"
	"log"
	"os"

	"github.com/chaosvermittlung/schaschlik/config"
)

//Messages is the channel for messages to be written to the file
var Messages chan string

//Path is the path to the file
var Path string

//Setup takes a config.Config and sets up the Message Writer
func Setup(conf config.Config) {
	Messages = make(chan string)
	Path = conf.Path
	go printRequest()
}

func printRequest() {
	for {
		m := <-Messages
		fmt.Println("Message", Messages)
		fmt.Println("path", Path)
		f, err := os.OpenFile(Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.WriteString(m)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}
