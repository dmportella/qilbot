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

// RespondToUser Sends a response to the user.
func (s *DiscordSession) RespondToUser(m *discordgo.MessageCreate, message string) {
	if s.onlyReplyWithDirectMessages {
		s.RespondToUserChannel(m, message)
	} else if s.onlyReplyInSpecificChannel {
		s.RespondInSpecificChannel(m, message, s.botChannelID)
	} else {
		s.RespondInChannel(m, message)
	}
}

// RespondInSpecificChannel Sends a response to specific channel regardless of where the commands was issued.
func (s *DiscordSession) RespondInSpecificChannel(m *discordgo.MessageCreate, message string, channelID string) {
	_, _ = s.ChannelMessageSend(channelID, message)
}

// RespondInChannel Sends a response to the channel where the commands was issued.
func (s *DiscordSession) RespondInChannel(m *discordgo.MessageCreate, message string) {
	_, _ = s.ChannelMessageSend(m.ChannelID, message)
}

// RespondToUserChannel Sends a response to the user channel as in a direct message.
func (s *DiscordSession) RespondToUserChannel(m *discordgo.MessageCreate, message string) {
	channel, _ := s.UserChannelCreate(m.Author.ID)

	_, _ = s.ChannelMessageSend(channel.ID, message)
}
