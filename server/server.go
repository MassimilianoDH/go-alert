package server

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"go-alert/bot"
	"go-alert/config"
	"go-alert/models"
	"log"
	"net/http"
	"text/template"
	"time"

	"gopkg.in/telebot.v3"
)

var alert models.Alert

type application struct {
	auth struct {
		username string
		password string
	}
}

func StartServer() {
	app := new(application)

	app.auth.username = config.Username
	app.auth.password = config.Password

	if app.auth.username == "" {
		log.Fatal("Server: Basic Auth Username must be provided!")
	}

	if app.auth.password == "" {
		log.Fatal("Server: Basic Auth Password must be provided!")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", app.basicAuth(app.webhook))

	srv := &http.Server{
		Addr:         config.Port,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server: Starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func (app *application) webhook(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&alert)

	if err != nil {
		log.Println("Server: Decoder: ", err)
		return
	}

	log.Println("Server: Incoming Alert from: ", alert.Incident.ScopingProjectID)

	if bot.ExpChat == 0 {
		log.Println("Server: Bot is off!")
		w.WriteHeader(500)
		return
	} else {

		// WIP
		data := models.Message{
			ProjectID:    alert.Incident.ScopingProjectID,
			ResourceType: alert.Incident.ResourceTypeDisplayName,
			PolicyName:   alert.Incident.PolicyName,
			ThreatLevel:  alert.Incident.PolicyUserLabels.UserLabel1,
			Summary:      alert.Incident.Summary,
			URL:          alert.Incident.URL,
		}

		t, err := template.ParseFiles("templates/message.txt")
		if err != nil {
			log.Fatal(err)
		}

		var b bytes.Buffer

		err = t.Execute(&b, &data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(b.String())
		// WIP

		log.Println("Server: Message sent to Telegram group: ", bot.ExpChat)
		bot.ExpBot.Send(bot.ExpChat, b.String(), telebot.ModeHTML)
		w.WriteHeader(200)
		return
	}

}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				log.Println("Server: Basic Auth successfull!")
				return
			}
		}
		log.Println("Server: Basic Auth failed!")
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
