package bot

import (
	"context"
	"go-alert/config"
	"log"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/discord"
	"github.com/nikoksr/notify/service/telegram"
)

// ExpBot will expose bot to server.go
var ExpBot *notify.Notify

func StartBot() {
	// Create a notify instance
	notify := notify.New()

	// Create a Discord service
	discordService := discord.New()
	discordService.AuthenticateWithBotToken(config.DiscordBotToken)
	discordService.AddReceivers(config.DiscordChatID)

	// Create a Telegram service
	telegramService, err := telegram.New(config.TelegramBotToken)
	telegramService.AddReceivers(config.TelegramChatID)

	if err != nil {
		log.Fatal(err)
	}

	// // Create a Slack service
	// slackService := slack.New("OAUTH_TOKEN")
	// slackService.AddReceivers("CHANNEL_ID")

	// Tell our notifier to use the multiple services.
	notify.UseServices(discordService, telegramService)

	// ExpBot will expose bot to server.go
	ExpBot = notify

	// Send message to confirm bot start
	err = notify.Send(
		context.Background(),
		"TOPIC: Bot Start",
		"I am a bot written in Go!",
	)

	if err != nil {
		log.Fatal(err)
	}
}
