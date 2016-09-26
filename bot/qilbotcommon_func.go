package bot

import (
	"bytes"
	"fmt"
	"github.com/dmportella/qilbot/logging"
	"strings"
)

func (qilbot *Qilbot) initialiseCommands() {

	helpCmd := QilbotCommand{
		Command:     "help",
		Template:    "!help or !help *command*",
		Description: "Display a list of commands available to qilbot and more information about specific commands.",
		Execute: func(ctx *QilbotCommandContext) {
			qilbot.helpCommand(ctx)
		},
	}

	setCmd := QilbotCommand{
		Command:     "set",
		Template:    "!set *variable* *value*",
		Description: "Changes the settings of qilbot at runtime.",
		Execute: func(ctx *QilbotCommandContext) {
			qilbot.setCommand(ctx)
		},
	}

	qilbot.AddCommand(&helpCmd)
	qilbot.AddCommand(&setCmd)
}

func (qilbot *Qilbot) helpCommand(ctx *QilbotCommandContext) {
	var buffer bytes.Buffer

	logging.Trace.Println("help comand text", ctx.CommandText)

	specificCommand := ctx.CommandText != ""

	for _, command := range qilbot.commands {
		if command.settings != nil && command.settings.Disabled {
			continue
		}

		if strings.Compare(strings.ToLower(ctx.CommandText), strings.ToLower(command.Command)) == 0 {
			buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))
		} else if !specificCommand {
			buffer.WriteString(fmt.Sprintf("**%s** (%s): %s\n", command.Command, command.Template, command.Description))
		}
	}

	ctx.RespondToUser(buffer.String())
}

func (qilbot *Qilbot) setCommand(ctx *QilbotCommandContext) {
	var buffer bytes.Buffer

	logging.Trace.Println("set comand text", ctx.CommandText)

	if ctx.IsOwnerOfGuild() {
		buffer.WriteString("This would have done something.")
	} else {
		buffer.WriteString("Only the Server owner can change the bot settings...")
	}

	ctx.RespondToUser(buffer.String())
}
