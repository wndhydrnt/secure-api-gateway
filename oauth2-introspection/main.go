package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	hardCodedToken = "abc123"
)

type introspectionResponse struct {
	Active    bool   `json:"active"`
	ClientID  string `json:"client_id"`
	ExpiresAt int64  `json:"exp"`
	Scope     string `json:"scope"`
	Subject   string `json:"sub"`
	Username  string `json:"username"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/oauth2/introspection" {
			err := r.ParseForm()
			if err != nil {
				log.Println("error parsing form: ", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			token := r.Form.Get("token")
			if token == "" {
				log.Println("token is empty")
				http.Error(w, "token is empty", http.StatusBadRequest)
				return
			}

			if token != hardCodedToken {
				log.Println("token is wrong")
				http.Error(w, "token is wrong", http.StatusBadRequest)
				return
			}

			b, err := json.Marshal(&introspectionResponse{
				Active:    true,
				ClientID:  "foobar",
				ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
				Scope:     "read",
				Subject:   "user-subject",
				Username:  "a-user",
			})
			if err != nil {
				log.Println("error encoding response: ", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		payload := fmt.Sprintf("method %s on route %s not supported", r.Method, r.URL.Path)
		log.Println(payload)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(payload))
		return
	})

	log.Fatal(http.ListenAndServe(":8484", nil))
}
