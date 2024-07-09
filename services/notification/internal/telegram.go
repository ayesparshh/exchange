package internal

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type Telegram struct {
	bot    *tgbotapi.BotAPI
	chatID int64
	pubsub PubSub
}

var (
	ConfigsCommand        = "configs"
	BalanceCommand        = "balance"
	PositionsCommand      = "positions"
	StatsCommand          = "stats"
	EnableTradingCommand  = "enable"
	DisableTradingCommand = "disable"
	DumpCommand           = "dump"
)

func NewTelegramBot(token string, chatId int64, pubsub PubSub) Telegram {
	log.Trace().Msg("Creating new telegram bot")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating telegram bot")
	}

	t := Telegram{bot, chatId, pubsub}
	t.SetDefaultCommands()

	return t
}

func (t *Telegram) SetDefaultCommands() {
	log.Trace().Msg("Setting default commands")
	configs := tgbotapi.BotCommand{Command: ConfigsCommand, Description: "Get configs"}
	balance := tgbotapi.BotCommand{Command: BalanceCommand, Description: "Get balance"}
	// positions := tgbotapi.BotCommand{PositionsCommand, "Get positions"}
	// stats := tgbotapi.BotCommand{StatsCommand, "Get statistics"}
	// enableTrading := tgbotapi.BotCommand{EnableTradingCommand, "Enable trading"}
	// disableTrading := tgbotapi.BotCommand{DisableTradingCommand, "Disable trading"}
	// dump := tgbotapi.BotCommand{DumpCommand, "Dump asset"}

	config := tgbotapi.NewSetMyCommands(configs, balance) //, positions, stats, enableTrading, disableTrading, dump)

	_, err := t.bot.Request(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error setting default commands")
	}
}

func (t *Telegram) ListenForCommands() {
	log.Trace().Msg("Listening for commands")

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates := t.bot.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// if update.Message.Chat.ID != t.chatID {
		// 	continue
		// }

		if !update.Message.IsCommand() {
			continue
		}

		args := update.Message.CommandArguments()
		message := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		message.ParseMode = tgbotapi.ModeHTML

		command := update.Message.Command()

		log.Info().Str("command", command).Str("args", args).Msg("Received command")

		switch command {
		case ConfigsCommand:
			var r GetConfigsResponse

			err := t.pubsub.Request(GetConfigsEvent, nil, &r)
			if err != nil {
				message.Text = err.Error()
				// t.bot.Send(message)
			} else {
				message.Text = t.FormatConfigsMessage(r)
			}
		case BalanceCommand:
			var r GetBalanceResponse

			err := t.pubsub.Request(GetBalanceEvent, nil, &r)

			if err != nil {
				message.Text = err.Error()
			} else {
				message.Text = t.FormatBalanceMessage(r)
			}
		// case StatsCommand:
		// 	var r GetStatsResponse

		// 	err := t.pubsub.Request(GetStatsEvent, nil, &r)

		// 	if err != nil {
		// 		message.Text = err.Error()
		// 	} else {
		// 		message.Text = t.FormatStatsMessage(r)
		// 	}
		default:
			message.Text = "Command not defined"
		}
		if update.Message.Chat.ID != t.chatID {
			from := update.Message.From

			log.Warn().Str("name", from.FirstName).Int("ID", int(from.ID)).Msg("Unauthorized Activity")
			message.Text = "You are not authorized, your activity has been recorded."

			notification := fmt.Sprintf("Unauthorized Activity\n\nID: %v\nName: %v", from.ID, from.FirstName)
			t.SendMessage(CriticalErrorEvent, notification)
		}

		_, err := t.bot.Send(message)
		if err != nil {
			log.Error().Err(err).Msg("TelegramBot.ListenForCommands")
		}
	}
}

func (t Telegram) SendMessage(event string, msg string) {
	log.Info().Str("event", event).Msg("TelegramBot.SendMessage")

	message := tgbotapi.NewMessage(t.chatID, msg)
	message.ParseMode = tgbotapi.ModeHTML

	_, err := t.bot.Send(message)

	if err != nil {
		log.Error().Err(err).Msg("TelegramBot.SendMessage")
	}
}
