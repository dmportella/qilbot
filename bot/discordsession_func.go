package bot

import (
	"github.com/bwmarrin/discordgo"
)

// IsOwnerOfGuild Check that the author of the message is the onwer of the guild.
func (s *DiscordSession) IsOwnerOfGuild(m *discordgo.MessageCreate) bool {
	channel, _ := s.Channel(m.ChannelID)

	guild, _ := s.Guild(channel.GuildID)

	return guild.OwnerID == m.Author.ID
}
