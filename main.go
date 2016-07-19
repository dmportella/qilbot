package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math"
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
	fmt.Printf("Golang tutorial version %s, branch %s at revision %s.\n\rDaniel Portella (c) 2016\n\r", Build, Branch, Revision)

	bava := [3]float64{83.40625, -134.3125, -80.75}
	sothis := [3]float64{-352.78125, 10.5, -346.34375}

	deltaX := bava[0] - sothis[0]
	deltaY := bava[1] - sothis[1]
	deltaZ := bava[2] - sothis[2]

	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)

	fmt.Printf("distance between bava and sothis is: %f\n", distance)

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
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
