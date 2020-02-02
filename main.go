package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/emaele/jackettbot/jackett"
)

var tBot *tgbotapi.BotAPI
var err error

func main() {
	// Telegram auth
	tBot, err = tgbotapi.NewBotAPI(telegramTokenBot)
	if err != nil {
		log.Panic(err)
	}

	// If the -dev flag is used set debug mode on telegram and mariadb client
	if dev {
		tBot.Debug = true
	}

	log.Printf("Authorized on account %s", tBot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tBot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	jclient := jackett.New(jackettAPIKey, jackettAddress)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Chat.IsPrivate() {
				go mainBot(jclient, *update.Message)
			}
		}
	}
}

func mainBot(jclient *jackett.Client, message tgbotapi.Message) {

	// Handle /cancel command
	if message.IsCommand() && message.Command() == "cancel" {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Search has been cancelled.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	// If text has the ⬇️ emoji as a prefix, handle it as a user choice for torrent
	if strings.HasPrefix(message.Text, "⬇️") {

		// Replace all spaces
		res, err := jclient.Search(strings.ReplaceAll(strings.TrimPrefix(message.Text, "⬇️ "), " ", "%20"))
		if err != nil {
			log.Println(err)
			return
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s\n\n`%s`", res.Results[0].Title, res.Results[0].MagnetURI))
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		msg.ParseMode = "Markdown"

		tBot.Send(msg)
		return
	}

	res, err := jclient.Search(strings.ReplaceAll(message.Text, " ", "%20"))
	if err != nil {
		log.Println(err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Select the torrent:")
	msg.ReplyMarkup = genKeyboard(res)

	tBot.Send(msg)
}

func genKeyboard(results jackett.SearchResults) tgbotapi.ReplyKeyboardMarkup {

	var buttons [][]tgbotapi.KeyboardButton

	cancelRow := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/cancel"))
	buttons = append(buttons, cancelRow)

	for index := 0; index < 20; index++ {
		if results.Results[index].MagnetURI == "" {
			continue
		}

		b := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(fmt.Sprintf("⬇️ %s", results.Results[index].Title)))
		buttons = append(buttons, b)
	}

	keyboard := tgbotapi.NewReplyKeyboard(buttons...)
	keyboard.OneTimeKeyboard = true

	return keyboard
}
