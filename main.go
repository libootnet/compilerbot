package main

import (
	"compilerbot/src"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BotToken string
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("env read error: %v\n", err)
		return
	}

	BotToken = os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	dg.AddHandler(src.MessageContent)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
