package bot

import (
	"io"
)

// IsOwnerOfGuild Check that the author of the message is the onwer of the guild.
func (ctx *QilbotCommandContext) IsOwnerOfGuild() bool {
	channel, _ := ctx.discordSession.Channel(ctx.Message.ChannelID)

	guild, _ := ctx.discordSession.Guild(channel.GuildID)

	return guild.OwnerID == ctx.Message.Author.ID
}

// RespondToUser Sends a response to the user.
func (ctx *QilbotCommandContext) RespondToUser(message string) {
	if ctx.command.settings.OnlyUsableOnDirectMessage {
		ctx.respondToUserChannel(message)
	} else if ctx.command.settings.OnlyUsableOnChannelID != "" {
		ctx.respondInSpecificChannel(message)
	} else {
		ctx.respondInChannel(message)
	}
}

func (ctx *QilbotCommandContext) respondInSpecificChannel(message string) {
	_, _ = ctx.discordSession.ChannelMessageSend(ctx.command.settings.OnlyUsableOnChannelID, message)
}

func (ctx *QilbotCommandContext) respondInChannel(message string) {
	_, _ = ctx.discordSession.ChannelMessageSend(ctx.Message.ChannelID, message)
}

func (ctx *QilbotCommandContext) respondToUserChannel(message string) {
	channel, _ := ctx.discordSession.UserChannelCreate(ctx.Message.Author.ID)

	_, _ = ctx.discordSession.ChannelMessageSend(channel.ID, message)
}

// ChannelTyping tells discord to display the typing action.
func (ctx *QilbotCommandContext) ChannelTyping() {
	if ctx.command.settings.OnlyUsableOnDirectMessage {
		channel, _ := ctx.discordSession.UserChannelCreate(ctx.Message.Author.ID)
		_ = ctx.discordSession.ChannelTyping(channel.ID)
	} else if ctx.command.settings.OnlyUsableOnChannelID != "" {
		_ = ctx.discordSession.ChannelTyping(ctx.command.settings.OnlyUsableOnChannelID)
	} else {
		_ = ctx.discordSession.ChannelTyping(ctx.Message.ChannelID)
	}

}

// RespondToUserWithFile sends a message to the user with an attachment.
func (ctx *QilbotCommandContext) RespondToUserWithFile(content string, name string, reader io.Reader) {
	if ctx.command.settings.OnlyUsableOnDirectMessage {
		channel, _ := ctx.discordSession.UserChannelCreate(ctx.Message.Author.ID)
		_, _ = ctx.discordSession.ChannelFileSendWithMessage(channel.ID, content, name, reader)
	} else if ctx.command.settings.OnlyUsableOnChannelID != "" {
		_, _ = ctx.discordSession.ChannelFileSendWithMessage(ctx.command.settings.OnlyUsableOnChannelID, content, name, reader)
	} else {
		_, _ = ctx.discordSession.ChannelFileSendWithMessage(ctx.Message.ChannelID, content, name, reader)
	}
}
