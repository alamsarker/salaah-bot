package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		"<!channel>Assalam, Please be prepared for the jamaah of %s salaah at %s ",
		salaahName, time,
	)

	payload := strings.NewReader(fmt.Sprintf(
		"token=%s&channel=%s&text=%s&username=%s",
		s.Token, s.Channel, message, s.UserName,
	))

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
