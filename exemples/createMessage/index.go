package main

import (
	"fmt"
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
							commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
								"test": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
									//autotrole.CreateMessage(s, i, content)
									
								},
								}
)

func main(){
	dg, err := discordgo.New("Bot " + os.Args[1])
	if err != nil { fmt.Println(err.Error())}

	err = dg.Open()
	if err != nil { fmt.Println(err.Error())}
	fmt.Println("Bot started")

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		}
		_, err := s.ApplicationCommandCreate("1045028748728217700", "1022982534516195398", &discordgo.ApplicationCommand{
			Name:        "test",
			Description: "Test the buttons if you got courage",
			})
		if err != nil { fmt.Println(err.Error())}

	})


	sc := make(chan os.Signal, 0)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	err = dg.Close()
	if err != nil { fmt.Println(err.Error())}
}

