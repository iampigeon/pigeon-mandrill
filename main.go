package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/keighl/mandrill"
)

type Message struct {
	Token     string `json:"token"`
	ToEmail   string `json:"email"`
	Type      string `json:"type"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Subject   string `json:"subject"`
	HTML      string `json:"html"`
	Text      string `json:"text"`
}

func main() {
	//TODO: use port defined on env values or specify at static url, example: http://pigeon-mandrill.wisegrowth.io
	http.ListenAndServe(":5152", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var msg Message

		//TODO: remove logger
		fmt.Println("here on ")

		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			// TODO use an error that better explains the problem
			http.Error(w, "bad request", http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		r.Body.Close()

		// TODO remove logger
		fmt.Printf("%#v\n", msg)

		client := m.ClientWithKey(msg.Token)

		email := &m.Message{}
		email.AddRecipient(msg.FromEmail, msg.FromName, msg.Type)
		email.FromEmail = msg.FromEmail
		email.FromName = msg.FromName
		email.Subject = msg.Subject
		email.HTML = msg.HTML
		email.Text = msg.Text

		response, err := client.MessagesSend(email)
		if err != nil {
			// TODO use an error that better explains the problem
			http.Error(w, "forbidden", http.StatusForbidden)
			fmt.Println(err)
		}

		// TODO remove Logger
		fmt.Println("message sent")
		fmt.Println(response)

		w.WriteHeader(200)
	}))
}
