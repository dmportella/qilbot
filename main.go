package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dmportella/qilbot/edsm"
	"regexp"
	"strings"
)

// Build version of the binary
var Build string

// Revision number of the binary
var Revision string

// Branch name of the binary
var Branch string

// Variables used for command line parameters
var (
	Email    string
	Password string
	Token    string
	BotID    string
)

func init() {

	flag.StringVar(&Email, "e", "", "Account Email")
	flag.StringVar(&Password, "p", "", "Account Password")
	flag.StringVar(&Token, "t", "", "Account Token")
	flag.Parse()
}

func main() {
	go startbot()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}

func startbot() {
	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New(Email, Password, Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
		panic(err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		panic(err)
	}
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	actionPattern := regexp.MustCompile(`^<@([0-9]+)>\s([a-z]+)\s?(.*)`)
	matches := actionPattern.FindStringSubmatch(m.Content)

	fmt.Println(m.Content)

	if len(matches) >= 3 && matches[1] == BotID {

		switch matches[2] {
		case "distance":
			placesPattern := regexp.MustCompile(`^(.*)\s?\/\s?(.*)`)
			placeMatches := placesPattern.FindStringSubmatch(matches[3])

			if len(placeMatches) != 0 {
				sys1 := strings.TrimSpace(placeMatches[1])
				sys2 := strings.TrimSpace(placeMatches[2])
				if distance, err := edsm.GetDistanceBetweenTwoSystems(sys1, sys2); err == nil {
					_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The distance between **%s** and **%s** is **%.2fly**.", placeMatches[1], placeMatches[2], distance))
				} else {
					_, _ = s.ChannelMessageSend(m.ChannelID, "There was an error trying to get the distance.")
				}
			} else {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Please give ma two places, format: distance **A** / **B**")
			}
			break
		case "ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
			break
		case "pong":
			_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
			break
		default:
			_, _ = s.ChannelMessageSend(m.ChannelID, "What?! I dont know what you mean...")
		}
	}

	//
}
