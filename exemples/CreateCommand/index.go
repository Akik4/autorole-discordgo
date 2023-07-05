package main

import (
	autotrole "github.com/akik4/autorole-discordgo/autorole"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	btn = []autotrole.ButtonFields{
		{
			Buttons: []autotrole.ButtonField{
				{
					Label:    "Test",
					Style:    0,
					Disabled: false,
					EmojiID:  "👍",
					Customid: "test",
					},
					{
					Label:    "Sauce",
					Style:    0,
					Disabled: false,
					EmojiID:  "👀",
					Customid: "sauce",
					},
					},
					},
					}
					content = []autotrole.MessageFormat{
						{
							ResponseType: discordgo.InteractionResponseChannelMessageWithSource,
							Content: "Voici un message basic",
							Buttons: btn,
							},
							}
)

func main(){
	dg, err := discordgo.New("Bot " + os.Args[1])
	if err != nil { return }

	dg.Open()

	autotrole.CreateCommand(dg, "test", content)

	sc := make(chan os.Signal, 0)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}