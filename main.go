package main

import (
	"github.com/alamsarker/salaah-bot/util"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sendMessage(m MessageInterface, salaahName, time string) {
	m.Send(salaahName, time)
}

func main() {

	juhr := os.Getenv("JUHR")
	asr := os.Getenv("ASR")
	magrib := os.Getenv("MAGRIB")

	s := gocron.NewScheduler()

	slack := Slack{
		Url:      os.Getenv("SLACK_URL"),
		Token:    os.Getenv("SLACK_TOKEN"),
		Channel:  os.Getenv("SLACK_CHANNEL"),
		UserName: os.Getenv("SLACK_USER_NAME"),
	}

	s.Every(1).Day().At(util.GetCronTime(juhr)).Do(sendMessage, slack, "Juhr", string(juhr))
	s.Every(1).Day().At(util.GetCronTime(asr)).Do(sendMessage, slack, "Asr", string(asr))
	s.Every(1).Day().At(util.GetCronTime(magrib)).Do(sendMessage, slack, "Magrib", string(magrib))

	<-s.Start()
}
