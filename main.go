package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func main() {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("can't get the api key")
	}
	botUser, exists := os.LookupEnv("BOT_USER")
	if !exists {
		log.Fatal("can't get the bot username")
	}
	keyword, exists := os.LookupEnv("KEY_WORD")
	if !exists {
		keyword = "clippy"
	}
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch msg.Data.(type) {
		case *slack.MessageEvent:
			message, ok := msg.Data.(*slack.MessageEvent)
			if !ok {
				log.Printf("couldn't cast message")
			}
			go handleMessage(keyword, botUser, api, message)
		default:
		}
	}
}

func sendReply(api *slack.Client, channel, botUser, message string) {
	params := slack.PostMessageParameters{
		Markdown:    true,
		Username:    botUser,
		UnfurlLinks: true,
	}
	_, _, err := api.PostMessage(channel, generateMessage(message), params)
	if err != nil {
		fmt.Printf("error sending reply %s", err)
	}
}

func handleMessage(keyword, botUser string, api *slack.Client, message *slack.MessageEvent) {
	if strings.HasPrefix(message.Text, keyword) {
		go sendReply(api, message.Channel, botUser, message.Text[len(keyword):])
	} else {
		fmt.Printf("message text %s doesn't have %s prefix\n", message.Text, keyword)
	}
}

func generateMessage(message string) string {

	var lines bytes.Buffer
	words := strings.Split(message, " ")
	//word, words := words[0], words[1:]
	for len(words) > 0 {
		var word string
		var line string
		for len(words) > 0 && len(line)+len(words[0])+1 < 39 {
			word, words = words[0], words[1:]
			line = line + " " + word

		}
		outputLine := fmt.Sprintf("   | %s", line)
		filler := 43 - len(outputLine)
		for i := 0; i < filler; i++ {
			outputLine = outputLine + " "
		}
		outputLine = outputLine + "|"
		lines.WriteString(outputLine)
		lines.WriteString("\n")
		line = ""
	}

	header := `
   -----------------------------------------
`

	footer := `   -----------------------------------------
	\
	 \
		__
	   /  \
	   |  |
	   @  @
	   |  |
	   || |/
	   || ||
	   |\_/|
	   \___/
	`

	return "```\n" + header + lines.String() + footer + "\n```"
}
