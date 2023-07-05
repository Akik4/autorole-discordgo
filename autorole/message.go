package autotrole

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)
//Some field can be ignored but at least customID is necessary
type ButtonField struct {
	Label string
	Style discordgo.ButtonStyle
	Disabled bool
	EmojiID string
	Customid string // Most important field
}
//need ButtonField
type ButtonFields struct {
	Buttons []ButtonField
}
//need ButtonFields
//Most used ResponseType is discordgo.InteractionResponseChannelMessageWithSource
type MessageFormat struct {
	ResponseType discordgo.InteractionResponseType
	Content string
	Buttons []ButtonFields
}

//This function create the response to a slash command no need to setup discordgo.Session.InteractionRespond
func ResponseSlashCommand(s *discordgo.Session,i *discordgo.InteractionCreate ,field []MessageFormat) {
	var (
		responseType discordgo.InteractionResponseType
		content string
		buttons []ButtonFields
	)
	for _, e := range field {
		responseType = e.ResponseType
		content = e.Content
		buttons = e.Buttons
	}
	var component []discordgo.MessageComponent
	for _, e := range buttons {
		for _, single := range e.Buttons {
			component = append(component, discordgo.Button{
				Label:    single.Label,
				Style:    single.Style,
				Disabled: single.Disabled,
				Emoji:    discordgo.ComponentEmoji{Name: single.EmojiID},
				CustomID: single.Customid,
				})
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: responseType,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: component,
				},
			},
		},
	})
	if err != nil { return }
}

func SetRoleOnClick(s *discordgo.Session, i *discordgo.InteractionCreate, roleID string){
	for _, roles := range i.Member.Roles {
		fmt.Println(roles)
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You received a new role",
			Flags: discordgo.MessageFlagsEphemeral,
		},

	})
	if err != nil { fmt.Println(err.Error()); return}
}