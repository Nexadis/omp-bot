package pkg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackageCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get idx - get product by idx\n"+
			"/new product name - create new product\n"+
			"/edit idx changed name - edit existed product\n"+
			"/delete idx - delete product by idx\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Help: error sending reply message to chat - %v", err)
	}
}
