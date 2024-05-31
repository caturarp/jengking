package main

import "time"

type WebhookPayload struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Repository Repository `json:"repository"`
	Pusher     Pusher     `json:"pusher"`
	Commits    []Commit   `json:"commits"`
	Sender     Sender     `json:"sender"`
}

type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	GitUrl   string `json:"git_url"`
}

type Pusher struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Email    string `json:"email"`
}

type Sender struct {
	Username string `json:"login"`
	Email    string `json:"email"`
}

type Commit struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Distinct  bool      `json:"distinct"`
	Url       string    `json:"url"`
}
