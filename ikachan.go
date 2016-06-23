package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Request struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	VcsURL          string      `json:"vcs_url"`
	BuildURL        string      `json:"build_url"`
	BuildNum        int         `json:"build_num"`
	Branch          string      `json:"branch"`
	VcsRevision     string      `json:"vcs_revision"`
	CommitterName   string      `json:"committer_name"`
	CommitterEmail  string      `json:"committer_email"`
	Subject         string      `json:"subject"`
	Body            string      `json:"body"`
	Why             string      `json:"why"`
	DontBuild       interface{} `json:"dont_build"`
	QueuedAt        time.Time   `json:"queued_at"`
	StartTime       time.Time   `json:"start_time"`
	StopTime        time.Time   `json:"stop_time"`
	BuildTimeMillis int         `json:"build_time_millis"`
	Username        string      `json:"username"`
	Reponame        string      `json:"reponame"`
	Lifecycle       string      `json:"lifecycle"`
	Outcome         string      `json:"outcome"`
	Status          string      `json:"status"`
	RetryOf         interface{} `json:"retry_of"`
}

func (self Payload) Header() string {
	if self.Status == "success" {
		return "(successful)"
	} else {
		return "(failed)"
	}
}

func main() {
	baseUrl := flag.String("ikachan", "", "ikachan url")
	listen := flag.String("listen", ":8080", "gateway listen")
	flag.Parse()
	if baseUrl == nil || *baseUrl == "" {
		log.Print(*baseUrl)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t Request
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("JSON parse error: %v", err)
			fmt.Fprintf(w, "%v", err)
			return
		}

		r.ParseForm()
		channel := r.FormValue("channel")
		if channel == "" {
			log.Print("Missing channel")
			fmt.Fprintf(w, "Missing channel\n")
			return
		}
		message_type := r.FormValue("message_type")
		if message_type == "" {
			message_type = "privmsg"
		}

		payload := t.Payload
		message := fmt.Sprintf("%s [%s] %s [%s/%s#%s]\n%s", payload.Header(), strings.ToUpper(payload.Status), payload.Subject, payload.Username, payload.Reponame, payload.Branch, payload.BuildURL)
		log.Print(message)

		resp, err := http.PostForm(*baseUrl+"/"+message_type,
			url.Values{"channel": {channel}, "message": {message}})
		if err != nil {
			log.Printf("Cannot send message to %s: %v", channel, err)
			fmt.Fprintf(w, "Cannot send message to %s: %v\n", channel, err)
			return
		}

		log.Printf("Sent message to %s(%s): %v", channel, message_type, resp)
		fmt.Fprintf(w, "OK\n")
	})

	log.Printf("Starting httpd: %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
