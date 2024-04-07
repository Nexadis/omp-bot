package pkg

import (
	"encoding/json"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/pkg"
)

const pageSize = 5

func (c *LogisticPackageCommander) List(inputMessage *tgbotapi.Message, offset int) {
	msg := c.Page(inputMessage.Chat.ID, offset, pageSize)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticPackageCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c *LogisticPackageCommander) Page(ChatID int64, offset int, size int) tgbotapi.MessageConfig {
	packages := c.packageService.List()
	end := offset + pageSize
	if end > len(packages) {
		end = len(packages)
	}
	msg := tgbotapi.NewMessage(ChatID, FormatList(packages[offset:end]))

	cbPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "package",
		CallbackName: "list",
	}
	next := cbPath
	nextData, _ := json.Marshal(CallbackListData{
		Offset: offset + size,
	})
	next.CallbackData = string(nextData)

	prevData, _ := json.Marshal(CallbackListData{
		Offset: offset - size,
	})

	prev := cbPath
	prev.CallbackData = string(prevData)

	buttons := []tgbotapi.InlineKeyboardButton{}
	if offset != 0 {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev page", prev.String()))
	}
	if offset+size < len(packages) {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", next.String()))
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	return msg
}

func FormatList(list []pkg.Package) string {
	if len(list) == 0 {
		return "List is empty"
	}
	outString := strings.Builder{}
	outString.WriteString("Packages:\n\n")
	for _, l := range list {
		outString.WriteString(l.String())
		outString.WriteRune('\n')
	}
	return outString.String()
}
