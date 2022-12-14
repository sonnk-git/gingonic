package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

var TIME_TO_REMIND = [...]int{
	1,
	2,
	3,
	5,
	10,
	30,
	60,
	120,
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minutes().Do(func() {
		sendNotificationRequest(1)
	})

	s.Every(2).Minutes().Do(func() {
		sendNotificationRequest(2)
	})

	s.Every(3).Minutes().Do(func() {
		sendNotificationRequest(3)
	})

	s.Every(5).Minutes().Do(func() {
		sendNotificationRequest(5)
	})

	s.Every(10).Minutes().Do(func() {
		sendNotificationRequest(10)
	})

	s.Every(30).Minutes().Do(func() {
		sendNotificationRequest(30)
	})

	s.Every(60).Minutes().Do(func() {
		sendNotificationRequest(60)
	})

	s.Every(120).Minutes().Do(func() {
		sendNotificationRequest(120)
	})

	s.StartBlocking()
}

func sendNotificationRequest(everyMinute int) {
	postBody, _ := json.Marshal(map[string]int{
		"everyMinute": everyMinute,
	})
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8080/api/v1/send-notification", "application/json", requestBody)

	fmt.Printf("%+v\n", resp)

	if err != nil {
		fmt.Printf("error when sendNotificationRequest:  %v", err)
	} else {
		fmt.Println("sendNotification")
	}
}
