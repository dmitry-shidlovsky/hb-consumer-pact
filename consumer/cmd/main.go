package main

import (
	"log"
	"net/url"

	client "github.com/dmitry-shidlovsky/TestPact/consumer"
)

func main() {
	u, _ := url.Parse("http://localhost:8080")
	client := &client.Client{
		BaseURL: u,
	}

	users, err := client.GetUser(10)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
