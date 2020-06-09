package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"

	"github.com/eze-kiel/irc-bot/commands"
	irc "github.com/thoj/go-ircevent"
)

const channel = "#home"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : irc-bot <SERVER:PORT>")
		os.Exit(1)
	}

	server := os.Args[1]
	ircnick1 := "UselessBot"
	irccon := irc.IRC(ircnick1, "UselessBot")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = false
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(channel)
		irccon.Privmsg(channel, "Guess who's back ( ͡° ͜ʖ ͡°)")
	})
	irccon.AddCallback("366", func(e *irc.Event) {})
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		chatWithBot(irccon, e)
	})
	err := irccon.Connect(server)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}

func chatWithBot(con *irc.Connection, e *irc.Event) {
	switch strings.ToLower(e.Message()) {
	case "hey bot":
		con.Privmsg(channel, "hey "+e.Nick)
	case "wakanda":
		con.Privmsg(channel, "forever")
	case "chuck fact pls":
		con.Privmsg(channel, commands.ChuckFact())
	case "help pls":
		con.Privmsg(channel, commands.Help())
	case "number fact pls":
		con.Privmsg(channel, commands.Trivia())
	case "year fact pls":
		con.Privmsg(channel, commands.Year())
	case "date fact pls":
		con.Privmsg(channel, commands.Date())
	case "math fact pls":
		con.Privmsg(channel, commands.Math())
	case "xkcd pls":
		con.Privmsg(channel, commands.Xkcd())
	case "yes or no pls":
		con.Privmsg(channel, commands.YesOrNo())
	case "idea pls":
		con.Privmsg(channel, commands.Idea())
	case "kanye pls":
		con.Privmsg(channel, commands.Kanye())
	case "motivation pls":
		con.Privmsg(channel, commands.Motivation())
	}

}
