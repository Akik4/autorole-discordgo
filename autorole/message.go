package autotrole

import (
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

type ButtonFields struct {
	Buttons []ButtonField
}
type MessageFormat struct {
	ResponseType discordgo.InteractionResponseType
	Content string
	Buttons []ButtonFields
}

func CreateMessage(s *discordgo.Session,i *discordgo.InteractionCreate ,field []MessageFormat) {
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
