package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/webhook", webHook)
	log.Println("Listening for webhooks on port 5555...")
	log.Println(http.ListenAndServe(":5555", nil))
}

func webHook(w http.ResponseWriter, r *http.Request) {
	var payload WebhookPayload
	// Parse JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Failed to parse payload", http.StatusBadRequest)
		log.Println("Error parsing payload:", err)
		return
	}

	// Validate the webhook event type
	if r.Header.Get("X-GitHub-Event") != "push" {
		http.Error(w, "Invalid event type", http.StatusBadRequest)
		log.Println("Invalid event type:", r.Header.Get("X-GitHub-Event"))
		return
	}
	fmt.Printf("%+v\n", payload)
}

func runBuildScript(branchName, repoName string) error {
	cmd := exec.Command("sh", "build.sh", branchName, repoName)
	logFile, err := os.OpenFile("build.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	cmd.Stdout = mw
	cmd.Stderr = mw

	return cmd.Run()
}
