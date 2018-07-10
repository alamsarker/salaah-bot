package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func sendMessage(m MessageInterface, salaahName, time string) {
	m.Send(salaahName, time)
}

func getCronTime(t string) string {
	hm := strings.Split(t, ":")
	h, _ := strconv.Atoi(hm[0])
	m, _ := strconv.Atoi(hm[1])
	fmt.Println(h, m)
	namazT := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		int(h),
		int(m),
		0,
		0,
		time.UTC,
	).Add(time.Minute * -10)

	return fmt.Sprintf("%d:%d", namazT.Hour(), namazT.Second())
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	juhr := os.Getenv("JUHR")
	asr := os.Getenv("ASR")
	magrib := os.Getenv("MAGRIB")

	fmt.Println("Juhr at ", getCronTime(juhr))
	fmt.Println("Asr at ", getCronTime(asr))
	fmt.Println("Magri at ", getCronTime(magrib))

	s := gocron.NewScheduler()

	slack := Slack{
		Url:      os.Getenv("SLACK_URL"),
		Token:    os.Getenv("SLACK_TOKEN"),
		Channel:  os.Getenv("SLACK_CHANNEL"),
		UserName: os.Getenv("SLACK_USER_NAME"),
	}

	//s.Every(5).Seconds().Do(sendMessage, slack, "Juhr", string(juhr))

	fmt.Println(getCronTime(juhr))
	fmt.Println(slack)

	//s.Every(1).Day().At(getCronTime(juhr)).Do(sendmessage, slack, "juhr", string(juhr))
	//s.Every(1).Day().At(getCronTime(asr)).Do(sendmessage, slack, "Asr", string(asr))
	//s.Every(1).Day().At(getCronTime(magrib)).Do(sendmessage, slack, "Magrib", string(magrib))

	<-s.Start()
}
