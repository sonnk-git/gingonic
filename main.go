package main

import (
	"encoding/json"
	"fmt"
	"gingonic/db"
	"gingonic/models"
	"gingonic/route"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/go-co-op/gocron"
	"os"
	"time"
)

func main() {
	r := route.SetupRouter()

	vapidPublicKey := os.Getenv("VAPID_PUBLIC_KEY")
	vapidPrivateKey := os.Getenv("VAPID_PRIVATE_KEY")
	s := gocron.NewScheduler(time.UTC)
	task := func() {
		fmt.Println("DO")
		id := "01GK8A6RHYZPYZJTW6XSAR2C37"
		var sub models.Subscription
		db.Orm.Find(&sub, "user_id = ?", id)

		// Decode subscription
		s := &webpush.Subscription{}
		json.Unmarshal([]byte(sub.Sub), s)

		// Send Notification
		_, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
			Subscriber:      "example@example.com", // Do not include "mailto:"
			VAPIDPublicKey:  vapidPublicKey,
			VAPIDPrivateKey: vapidPrivateKey,
			TTL:             30,
		})
		if err != nil {
			fmt.Printf("error!")
		}
	}

	_, err :=s.Every(5).Seconds().Do(task)

	if err != nil {
		fmt.Println("Error run task")
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
