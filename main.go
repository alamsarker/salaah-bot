package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type MessageInterface interface {
	Send(salaahName, time string)
}

type Slack struct {
	Url, Token, Channel, UserName string
}

func (s Slack) Send(salaahName, time string) {

	message := fmt.Sprintf(
		"<!channel>Assalam, Please be prepared for %s salaah at %s ",
		salaahName, time,
	)

	payload := strings.NewReader(fmt.Sprintf(
		"token=%s&channel=%s&text=%s&username=%s",
		s.Token, s.Channel, message, s.UserName,
	))

	//fmt.Println(message, reqUrl)
	req, _ := http.NewRequest("POST", s.Url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(s)
	fmt.Println(res)
	fmt.Println(string(body))
}

func sendMessage(m MessageInterface, salaahName, time string) {
	m.Send(salaahName, time)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	juhr := os.Getenv("JUHR")
	asr := os.Getenv("ASR")
	magrib := os.Getenv("MAGRIB")

	fmt.Println("Juhr at ", juhr)
	fmt.Println("asr at ", asr)
	fmt.Println("magrib at ", magrib)

	s := gocron.NewScheduler()

	slack := Slack{
		Url:      os.Getenv("SLACK_URL"),
		Token:    os.Getenv("SLACK_TOKEN"),
		Channel:  os.Getenv("SLACK_CHANNEL"),
		UserName: os.Getenv("SLACK_USER_NAME"),
	}

	//s.Every(1).Seconds().Do(taskWithParam, 1, 2)
	s.Every(5).Seconds().Do(sendMessage, slack, "Juhr", string(juhr))
	<-s.Start()
}
