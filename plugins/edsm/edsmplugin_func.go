package edsm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
	"strconv"
	"strings"
)

const (
	name        = "Qilbot EDSM Plugin"
	description = "Client plugin for Elite Dangerous Star Map web site."
)

// NewPlugin creates a new instance of EDSMPlugin.
func NewPlugin(qilbot *bot.Qilbot) (plugin *Plugin) {
	debugMode := qilbot.InDebugMode()

	plugin = &Plugin{
		api: NewAPIClient(debugMode, "Discord Bot (https://github.com/dmportella/qilbot, 0.0.0)"),
	}

	distance := bot.QilbotCommand{
		Command:     "distance",
		Template:    "distance **sys1** / **sys2**",
		Description: "Uses the coords in EDSM to calculate the distance between the two star systems.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.distanceCommand(ctx)
		},
	}

	sphere := bot.QilbotCommand{
		Command:     "sphere",
		Template:    "sphere **sys1** 14.33ly",
		Description: "Returns a list of systems within a specified distance to specified system.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.sphereCommand(ctx)
		},
	}

	locate := bot.QilbotCommand{
		Command:     "locate",
		Template:    "locate **Commander Name**",
		Description: "Returns the location of a commander in EDSM.",
		Execute: func(ctx *bot.QilbotCommandContext) {
			plugin.locateCommand(ctx)
		},
	}

	qilbot.AddCommand(&distance)
	qilbot.AddCommand(&sphere)
	qilbot.AddCommand(&locate)

	return
}

func (plugin *Plugin) Name() string {
	return name
}

func (plugin *Plugin) Description() string {
	return description
}

func (plugin *Plugin) locateCommand(ctx *bot.QilbotCommandContext) {
	var buffer bytes.Buffer

	if cmdrPos, ok1 := plugin.api.GetPosition(ctx.CommandText); ok1 == nil && cmdrPos.MSGNum == 100 {
		logging.Trace.Println(fmt.Sprintf("%#v", cmdrPos))

		var header string

		if cmdrPos.System == "" {
			header = fmt.Sprintf("Player was found but he or she may not be sharing their location publicly.\r\nThe commander in question should check their settings in EDSM.")
		} else {
			header = fmt.Sprintf("Player is currently at **%s**", cmdrPos.System)
		}

		buffer.WriteString(header)

		ctx.RespondToUser(buffer.String())
	} else {
		buffer.WriteString("Player not found, the commander doesn't exist or they have not registered with EDSM.")

		ctx.RespondToUser(buffer.String())
	}
}

func (plugin *Plugin) sphereCommand(ctx *bot.QilbotCommandContext) {
	placeMatches := regexMatchSphereCommand(ctx.CommandText)

	logging.Trace.Println(placeMatches)

	if len(placeMatches) >= 3 {
		systemName := strings.TrimSpace(placeMatches[1])
		distance := strings.TrimSpace(placeMatches[2])

		logging.Trace.Println("systemname", systemName, "distance", distance)

		ctx.ChannelTyping()

		if sys1, ok1 := plugin.api.GetSystem(systemName); ok1 == nil {
			if systems, ok2 := plugin.getSphereSystems(sys1.Name, distance); ok2 == nil {
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

					ctx.RespondToUser("Response is too large for discord please narrow you search.")
				} else if buffer.Len() > 2000 {
					logging.Trace.Println("msg attach", buffer.Len())

					reader := bytes.NewReader(buffer.Bytes())

					ctx.RespondToUserWithFile(header, "Results.txt", reader)
				} else {
					logging.Trace.Println("msg oke", buffer.Len())

					ctx.RespondToUser(buffer.String())
				}
			} else {
				ctx.RespondToUser(ok2.Error())
			}
		} else {
			ctx.RespondToUser(ok1.Error())
		}
	}
}

func (plugin *Plugin) distanceCommand(ctx *bot.QilbotCommandContext) {
	placeMatches := regexMatchDistanceCommand(ctx.CommandText)

	if len(placeMatches) >= 3 {
		sys1 := strings.TrimSpace(placeMatches[1])
		sys2 := strings.TrimSpace(placeMatches[2])

		ctx.ChannelTyping()

		if distance, err := plugin.getDistanceBetweenTwoSystems(sys1, sys2); err == nil {
			ctx.RespondToUser(fmt.Sprintf("The distance between **%s** and **%s** is **%.2fly**.", sys1, sys2, distance))
		} else {
			logging.Trace.Println(err)

			ctx.RespondToUser(err.Error())
		}
	} else {
		ctx.RespondToUser("Please give me two places, format: distance **A** / **B**")
	}
}

func (plugin *Plugin) getDistanceBetweenTwoSystems(systemName1 string, systemName2 string) (distance float64, err error) {
	if sys1, ok1 := plugin.api.GetSystem(systemName1); ok1 == nil {
		if sys2, ok2 := plugin.api.GetSystem(systemName2); ok2 == nil {
			distance = calculateDistance(sys1.Coords, sys2.Coords)
		} else {
			logging.Trace.Println(ok2)
			err = fmt.Errorf("EDSM couldn't fetch information about **%s** system, please check the spelling", systemName2)
		}
	} else {
		logging.Trace.Println(ok1)
		err = fmt.Errorf("EDSM couldn't fetch information about **%s** system, please check the spelling", systemName1)
	}
	return
}

func (plugin *Plugin) getSphereSystems(systemName1 string, distance string) (systems []System, err error) {
	if value, ok1 := strconv.ParseFloat(distance, 64); ok1 == nil {
		if sysList, ok2 := plugin.api.GetSphereSystems(systemName1, value); ok2 == nil {
			systems = sysList
		} else {
			logging.Trace.Println(ok2)
			err = fmt.Errorf("EDSM couldn't fetch information about **%s** system, please check the spelling", systemName1)
		}
	} else {
		logging.Trace.Println(ok1)
		err = errors.New("The distance provided is not a float number")
	}
	return
}
