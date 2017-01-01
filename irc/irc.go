package irc

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/chaosvermittlung/schaschlik/config"
	"github.com/chaosvermittlung/schaschlik/writer"
	"github.com/quiteawful/qairc"
)

//Setup takes all IRC connections an creates one client for every connection
func Setup(conns []config.IRCConn) {
	for _, c := range conns {
		go createClient(c)
	}
}

func createClient(conn config.IRCConn) {
	client := qairc.QAIrc(conn.Nick, conn.Username)
	client.Address = conn.Adress
	client.UseTLS = conn.TLS
	client.TLSCfg = &tls.Config{InsecureSkipVerify: conn.SkipVerify}
	err := client.Run()
	log.Println(qairc.Numerics["005"])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("go!")
	for {
		m, status := <-client.Out
		if !status {
			log.Println("Out closed, exiting")
			client.Reconnect()
		}

		if m.Type == "001" {
			client.Join(conn.Channel)
		}
		if m.Type == "PRIVMSG" {
			l := len(m.Args)
			nick := m.Sender.Nick
			channel := m.Args[0 : l-1]
			msg := m.Args[l-1]
			time := m.Timestamp.Format(time.RFC3339)
			message := "[" + channel[0] + "]" + "[" + time + "]" + "[" + nick + "]:" + msg
			writer.Messages <- message
		}
	}
}
