package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loding env")
	}

	tPass := os.Getenv("TWITCH_PASS")

	client := twitch.NewClient("opti_21", tPass)

	client.OnConnect(func() {
		fmt.Println("Connected to twitch")
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	client.Join("opti_21")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}