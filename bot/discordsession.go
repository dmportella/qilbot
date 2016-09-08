package bot

import (
	"github.com/bwmarrin/discordgo"
)

// DiscordSession Discord wrapper for discordgo.
// This allows us to decorate inject methods into discordgo.Session object.
type DiscordSession struct {
	*discordgo.Session
}
