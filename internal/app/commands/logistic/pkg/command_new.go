package pkg

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackageCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	err := c.packageService.New(args)
	if err != nil {
		log.Printf("fail to create product %s: %v", args, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("product '%s' created", args),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
