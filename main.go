package main

import (
	"flag"
	"fmt"
	"github.com/dmportella/qilbot/bot"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/plugins/common"
	"github.com/dmportella/qilbot/plugins/edsm"
	"io/ioutil"
	"os"
)

// Set on build
var (
	Build    string
	Branch   string
	Revision string
)

// Variables used for command line parameters
var (
	Email    string
	Password string
	Token    string
	BotID    string
	Version  bool
	Verbose  bool
)

var (
	botInstance bot.Qilbot
)

func init() {
	const (
		defaultEmail    = ""
		emailUsage      = "The email of te discord user. Not required if -bot-token is provided."
		defaultPassword = ""
		passwordUsage   = "The password of te discord user. Not required if -bot-token is provided."
		defaultToken    = ""
		tokenUsage      = "The token for the dicord bot. For more information please visit: https://discordapp.com/developers"
	)

	flag.StringVar(&Email, "user-email", defaultEmail, emailUsage)
	flag.StringVar(&Password, "user-password", defaultPassword, passwordUsage)
	flag.StringVar(&Token, "bot-token", defaultToken, tokenUsage)
	flag.StringVar(&Email, "e", defaultEmail, emailUsage)
	flag.StringVar(&Password, "p", defaultPassword, passwordUsage)
	flag.StringVar(&Token, "t", defaultToken, tokenUsage)

	const (
		defaultVerbose = false
		verboseUsage   = "Redirect trace information to the standard out."
	)

	flag.BoolVar(&Verbose, "verbose", defaultVerbose, verboseUsage)
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	fmt.Printf("Qilbot - Version: %s Branch: %s Revision: %s.\n\rDaniel Portella (c) 2016\n\r", Build, Branch, Revision)

	if Verbose {
		logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	} else {
		logging.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	if Password == "" && Email == "" && Token == "" {
		logging.Error.Println("Please provide credentials.")
		os.Exit(1)
	}

	botConfig := bot.QilbotConfig{
		Email:    Email,
		Password: Password,
		Token:    Token,
		Debug:    Verbose,
	}

	bot, ok := bot.New(&botConfig)

	if ok != nil {
		os.Exit(2)
	} else {
		botInstance = *bot
	}

	loadPlugins()

	go startbot()

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}

func loadPlugins() {
	commonPlugin := common.NewPlugin(&botInstance)

	logging.Info.Println(commonPlugin.GetHelpText())

	edsmPlugin := edsm.NewPlugin(&botInstance)

	logging.Info.Println(edsmPlugin.GetHelpText())
}

func startbot() {
	if ok := botInstance.Start(); ok == nil {
		logging.Info.Println("Bot is now running.  Press CTRL-C to exit.")
	} else {
		panic(ok)
	}
}
