package bot

import (
	"go-alert/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

// ExpBot will expose bot to server.go
var ExpBot *tele.Bot

// ExpChat will expose chatID to be used as Recipient by server.go
var ExpChat tele.ChatID

func StartBot() {

	var err error

	adm := tele.User{
		ID: config.AdminID,
	}

	grp := tele.Chat{
		ID: config.GroupID,
	}

	pref := tele.Settings{
		Token:  config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	if adm.ID == 0 {
		log.Fatal("Bot: Telegram Admin ID must be provided!")
	}

	if grp.ID == 0 {
		log.Fatal("Bot: Telegram Group ID must be provided!")
	}

	if pref.Token == "" {
		log.Fatal("Bot: Telegram Bot Token must be provided!")
	}

	bot, err := tele.NewBot(pref)

	// ExpBot will expose bot to server.go

	ExpBot = bot

	if err != nil {
		log.Fatal("Bot: Error: ", err)
		return
	}

	// Ping/Pong

	bot.Handle("/ping", func(c tele.Context) error {

		user := c.Sender()

		log.Println("Bot: Ping received from user: ", user.ID)
		c.Send("Pong")

		return nil
	})

	// Added to wrong group
	// Bot can only be added to specified Telegram group

	bot.Handle(tele.OnAddedToGroup, func(c tele.Context) error {

		chat := c.Chat()
		var msg string

		if chat.ID == grp.ID {
			msg = "Access granted to " + chat.Title + "!"
			log.Println("Bot: Bot added to group: ", chat.ID)
			c.Send(msg)
		} else {
			msg = "Error: Bot can only be added to specified Telegram group."
			log.Println("Bot: Failed group add attempt from group: ", chat.ID)
			c.Send(msg)
			bot.Leave(chat)
		}

		return nil
	})

	// start
	// ExpChat will set to chatID

	bot.Handle("/start", func(c tele.Context) error {

		user := c.Sender()
		chat := c.Chat()
		var msg string

		if chat.ID == grp.ID {
			if user.ID == adm.ID {
				msg = "Access granted to " + user.Username + "!"
				ExpChat = tele.ChatID(grp.ID)
				log.Println("Bot: Bot started in group: ", chat.ID)
				c.Send(msg)
			} else {
				msg = "Error: Bot can only be started by specified Telegram admin."
				log.Println("Bot: Failed /start attempt from user: ", user.ID)
				c.Send(msg)
			}
		} else {
			msg = "Error: Bot can only be started in specified Telegram group."
			log.Println("Bot: Failed /start attempt from user: ", user.ID)
			c.Send(msg)
		}

		return nil
	})

	// stop
	// ExpChat will set to nil

	bot.Handle("/stop", func(c tele.Context) error {

		user := c.Sender()
		chat := c.Chat()
		var msg string

		if chat.ID == grp.ID {
			if user.ID == adm.ID {
				msg = "Access granted to " + user.Username + "!"
				ExpChat = 0
				log.Println("Bot: Bot stopped in group: ", chat.ID)
				c.Send(msg)
			} else {
				msg = "Error: Bot can only be stopped by specified Telegram admin."
				log.Println("Bot: Failed /stop attempt from user: ", user.ID)
				c.Send(msg)
			}
		} else {
			msg = "Error: Bot can only be stopped in specified Telegram group."
			log.Println("Bot: Failed /stop attempt from user: ", user.ID)
			c.Send(msg)
		}

		return nil
	})

	log.Println("Bot: Started successfully!")
	bot.Start()
}
