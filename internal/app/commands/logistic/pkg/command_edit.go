package pkg

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticPackageCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	vals := strings.SplitN(args, " ", 2)
	if len(args) != 2 {
		log.Println("wrong args", args)
		return
	}
	i, value := vals[0], vals[1]

	idx, err := strconv.Atoi(i)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	err = c.packageService.Edit(idx, value)
	if err != nil {
		log.Printf("fail to edit product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("product %d edited: '%s'", idx, value),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.Get: error sending reply message to chat - %v", err)
	}
}
