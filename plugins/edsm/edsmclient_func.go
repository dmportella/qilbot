package edsm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/utilities"
	"strconv"
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
				bot.CommandInformation{
					Command:     "sphere",
					Template:    "sphere **sys1** 14.33ly",
					Description: "Returns a list of systems within a specified distance to specified system.",
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
		} else {
			logging.Trace.Println(ok2)
		}
	} else {
		logging.Trace.Println(ok1)
	}

	err = errors.New("Unable to get distance between these systems.")

	return
}

func (self *EDSMPlugin) GetSphereSystems(systemName1 string, distance string) (systems []System, err error) {
	if value, ok1 := strconv.ParseFloat(distance, 64); ok1 == nil {
		if sysList, ok2 := getSphereSystems(systemName1, value); ok2 == nil {
			systems = sysList
			return
		} else {
			logging.Trace.Println(ok2)
		}
	} else {
		logging.Trace.Println(ok1)
	}

	err = errors.New("Unable to get nearest systems.")
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
			self.displayDistance(s, m, matches[3])
			break
		case "sphere":
			self.displaySphere(s, m, matches[3])
		default:
			return
		}
	}
}

func (self *EDSMPlugin) displaySphere(s *discordgo.Session, m *discordgo.MessageCreate, commandText string) {
	placeMatches := RegexMatchSphereCommand(commandText)

	logging.Trace.Println(placeMatches)

	if len(placeMatches) >= 3 {
		systemName := strings.TrimSpace(placeMatches[1])
		distance := strings.TrimSpace(placeMatches[2])

		logging.Trace.Println("systemname", systemName, "distance", distance)

		s.ChannelTyping(m.ChannelID)

		if sys1, ok1 := getSystem(systemName); ok1 == nil {
			if systems, ok2 := self.GetSphereSystems(sys1.Name, distance); ok2 == nil {
				var buffer bytes.Buffer

				header := fmt.Sprintf("Found **%d** systems within **%sly** of **%s**.\r\n", len(systems)-1, distance, sys1.Name)

				buffer.WriteString(header)

				buffer.Write([]byte("```\r\n"))

				for _, sys2 := range systems {
					if sys2.Name == sys1.Name {
						continue
					}

					fmt.Fprintf(&buffer, "%-30s\t\t\t\t%9.2fly\r\n", sys2.Name, calculateDistance(sys1.Coords, sys2.Coords))
				}

				buffer.Write([]byte("```\r\n"))

				if buffer.Len() > 8388608 {
					logging.Trace.Println("msg large", buffer.Len())

					_, _ = s.ChannelMessageSend(m.ChannelID, "Response is too large for discord please narrow you search.")
				} else if buffer.Len() > 2000 {
					logging.Trace.Println("msg attach", buffer.Len())
					reader := bytes.NewReader(buffer.Bytes())
					_, _ = s.ChannelFileSendWithMessage(m.ChannelID, header, "Results.txt", reader)
				} else {
					logging.Trace.Println("msg oke", buffer.Len())
					_, _ = s.ChannelMessageSend(m.ChannelID, buffer.String())
				}
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, ok2.Error())
			}
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, ok1.Error())
		}
	}
}

func (self *EDSMPlugin) displayDistance(s *discordgo.Session, m *discordgo.MessageCreate, commandText string) {
	placeMatches := RegexMatchDistanceCommand(commandText)

	if len(placeMatches) >= 3 {
		sys1 := strings.TrimSpace(placeMatches[1])
		sys2 := strings.TrimSpace(placeMatches[2])

		s.ChannelTyping(m.ChannelID)

		if distance, err := self.GetDistanceBetweenTwoSystems(sys1, sys2); err == nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The distance between **%s** and **%s** is **%.2fly**.", sys1, sys2, distance))
		} else {
			logging.Warning.Println(err)
			_, _ = s.ChannelMessageSend(m.ChannelID, "There was an error trying to get the distance.")
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Please give me two places, format: distance **A** / **B**")
	}
}
