package edsm

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/utilities"
	"strings"
)

const (
	NAME        = "EDSM Plugin"
	DESCRIPTION = "Client plugin for Elite Dangerous Star Map web site."
)

func New() EDSMPlugin {
	return EDSMPlugin{
		bot.Plugin{
			Name:        NAME,
			Description: DESCRIPTION,
			Commands: []bot.CommandInformation{
				bot.CommandInformation{
					Command:     "distance",
					Template:    "distance **sys1** / **sys2**",
					Description: "Uses the coords in EDSM to calculate the distance between the two star systems.",
				},
			},
		},
	}
}

func (self *EDSMPlugin) GetDistanceBetweenTwoSystems(systemName1 string, systemName2 string) (distance float64, err error) {
	if sys1, ok1 := getSystem(systemName1); ok1 == nil {
		if sys2, ok2 := getSystem(systemName2); ok2 == nil {
			distance = calculateDistance(sys1.Coords, sys2.Coords)
			return
		}
	}

	err = errors.New("Unable to get distance between these systems.")

	return
}

func (self *EDSMPlugin) Initialize(qilbot *bot.Qilbot) {
	self.Qilbot = qilbot
	qilbot.AddHandler(self.messageCreate)
}

func (self *EDSMPlugin) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if self.Plugin.Qilbot.IsBot(m.Author.ID) {
		return
	}

	matches := utilities.RegexMatchBotCommand(m.Content)

	if len(matches) >= 3 && self.Plugin.Qilbot.IsBot(matches[1]) {
		switch matches[2] {
		case "distance":
			placeMatches := RegexMatchDistanceCommand(matches[3])

			if len(placeMatches) >= 3 {
				sys1 := strings.TrimSpace(placeMatches[1])
				sys2 := strings.TrimSpace(placeMatches[2])
				if distance, err := self.GetDistanceBetweenTwoSystems(sys1, sys2); err == nil {
					_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The distance between **%s** and **%s** is **%.2fly**.", sys1, sys2, distance))
				} else {
					logging.Warning.Println(err)
					_, _ = s.ChannelMessageSend(m.ChannelID, "There was an error trying to get the distance.")
				}
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Please give me two places, format: distance **A** / **B**")
			}
			break
		default:
			return
		}
	}
}
