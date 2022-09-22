package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	dg.Identify.Intents |= discordgo.IntentMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Println("The message is")
	fmt.Println(m.Content)

	// If the message is "ping" reply with "Pong!"
	tiltMatched, err := regexp.MatchString(`.tilt`, m.Content)

	if err != nil {
		fmt.Println("Error regexp parsing message:", m.Content, err)
		return
	}

	tiltString := "Dude FUCK THIS YEAR. Honestly garbage. Playing the fucking highest scorer every fucking week but still putting up points while the rest of the shit teams get wins. Fuck this dude. I'll auto draft next year. Full fuckin tilt."

	if tiltMatched {
		s.ChannelMessageSend(m.ChannelID, tiltString)
	}

}
