package config

import (
	"os"
	"strconv"
)

// SERVER

// Server Port
var Port string = ":" + os.Getenv("SERVER_PORT")

// Basic Auth
var Username string = os.Getenv("AUTH_USERNAME")
var Password string = os.Getenv("AUTH_PASSWORD")

// Templates
var GCPTemplate string = os.Getenv("GCP_TEMPLATE")
var AZRTemplate string = os.Getenv("AZR_TEMPLATE")
var AWSTemplate string = os.Getenv("AWS_TEMPLATE")

// CHATS

// Discord
var DiscordBotToken string = os.Getenv("DISCORD_BOT_TOKEN")
var DiscordChatID string = os.Getenv("DISCORD_CHAT_ID")

// MSTEams
var TeamsWebhook string = os.Getenv("MSTEAMS_WEBHOOK")

// Slack
var SlackBotToken string = os.Getenv("SLACK_BOT_TOKEN")
var SlackChatID string = os.Getenv("SLACK_CHAT_ID")

// Telegram
var TelegramBotToken string = os.Getenv("TELEGRAM_BOT_TOKEN")
var TelegramChatID, _ = strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
