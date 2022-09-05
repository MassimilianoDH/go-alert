package server

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"go-alert/bot"
	"go-alert/config"
	"go-alert/models"
	"log"
	"net/http"
	"time"
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

		//
		//
		//
		id := "IncidentID:" + " " + alert.Incident.IncidentID + "\n" // FIX!
		sum := "Summary:" + " " + alert.Incident.Summary + "\n"
		url := "URL:" + " " + alert.Incident.URL + "\n"
		msg := id + sum + url
		//
		//
		//

		log.Println("Server: Message sent to Telegram group: ", bot.ExpChat)
		bot.ExpBot.Send(bot.ExpChat, msg)
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
