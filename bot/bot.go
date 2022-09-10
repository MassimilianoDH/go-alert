package bot

import (
	"context"
	"go-alert/config"
	"log"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/discord"
	"github.com/nikoksr/notify/service/msteams"
	"github.com/nikoksr/notify/service/slack"
	"github.com/nikoksr/notify/service/telegram"
)

// ExpBot will expose bot to server.go
var ExpBot *notify.Notify

func StartBot() {

	var err error

	// Create a notify instance
	notify := notify.New()

	// Create a Discord service
	if config.DiscordBotToken != "" && config.DiscordChatID != "" {
		discordService := discord.New()
		discordService.AuthenticateWithBotToken(config.DiscordBotToken)
		discordService.AddReceivers(config.DiscordChatID)
		notify.UseServices(discordService)
	}

	// Create a Microsoft Teams service
	if config.TeamsWebhook != "" {
		msTeamsService := msteams.New()
		msTeamsService.AddReceivers(config.TeamsWebhook)
		notify.UseServices(msTeamsService)
	}

	// Create a Slack service
	if config.SlackBotToken != "" && config.SlackChatID != "" {
		slackService := slack.New(config.SlackBotToken)
		slackService.AddReceivers(config.SlackChatID)
		notify.UseServices(slackService)
	}

	// Create a Telegram service
	if config.TelegramBotToken != "" && config.TelegramChatID != 0 {
		telegramService, err := telegram.New(config.TelegramBotToken)
		if err != nil {
			log.Fatal(err)
		}
		telegramService.SetParseMode(telegram.ModeMarkdown)
		telegramService.AddReceivers(config.TelegramChatID)
		notify.UseServices(telegramService)
	}

	// ExpBot will expose bot to server.go
	ExpBot = notify

	// Send message to confirm bot start
	err = notify.Send(
		context.Background(),
		"🤖🤖🤖  *Bot Start*  🤖🤖🤖",
		"Bot started successfully! ✅",
	)
	if err != nil {
		log.Fatal(err)
	}
}
