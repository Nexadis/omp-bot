package pkg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/pkg"
)

type LogisticPackageCommander struct {
	bot            *tgbotapi.BotAPI
	packageService *pkg.PackageService
}

func NewLogisticPackageCommander(
	bot *tgbotapi.BotAPI,
) *LogisticPackageCommander {
	packageService := pkg.NewPackageService()

	return &LogisticPackageCommander{
		bot:            bot,
		packageService: packageService,
	}
}

func (c *LogisticPackageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LogisticPackageCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LogisticPackageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg, 0)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
