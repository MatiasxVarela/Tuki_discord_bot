package main

import (
	"bot/discord/messages"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	TOKEN := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messages.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot iniciado")
}
