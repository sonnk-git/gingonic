package controllers

import (
	"encoding/json"
	"fmt"
	"gingonic/db"
	"gingonic/models"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SubscribeNotificationRequest struct {
	Subscription string
}

type SetSubscribeRequest struct {
	State bool
}

func GetInfo(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	email, _ := ctx.Get("email")
	username, _ := ctx.Get("username")
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":   true,
		"id":       id,
		"email":    email,
		"username": username,
	})
}

func SubscribeNotification(ctx *gin.Context) {
	var requestBody SubscribeNotificationRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}

	id, _ := ctx.Get("id")
	var subscription models.Subscription
	tx := db.Orm.Find(&subscription, "user_id = ?", id)
	if tx.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}

	subscription.Sub = requestBody.Subscription
	if tx.RowsAffected > 0 {
		db.Orm.Save(&subscription)
	} else {
		subscription.UserID = id.(string)
		subscription.CourseID = "cdcdscsd"
		subscription.EveryMinute = 1
		db.Orm.Create(&subscription)
	}

	fmt.Printf("%+v\n", subscription)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":       true,
		"message":      "subscribe notification for user " + id.(string),
		"subscription": subscription,
	})
}

func SendNotification(ctx *gin.Context) {
	vapidPublicKey := os.Getenv("VAPID_PUBLIC_KEY")
	vapidPrivateKey := os.Getenv("VAPID_PRIVATE_KEY")

	id, _ := ctx.Get("id")
	var sub models.Subscription
	tx := db.Orm.Find(&sub, "user_id = ?", id)
	if tx.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}

	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(sub.Sub), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "example@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusServiceUnavailable, gin.H{
			"status": false,
		})
	}
	defer resp.Body.Close()

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func SetSubscribe(ctx *gin.Context) {
	var req SetSubscribeRequest
	id, _ := ctx.Get("id")
	var sub models.Subscription

	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}

	tx := db.Orm.Find(&sub, "user_id = ?", id)
	if tx.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": tx.Error.Error(),
		})
		return
	}

	if tx.RowsAffected > 0 {
		sub.SubscribeState = req.State
		db.Orm.Save(&sub)
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "cannot get subscription, please subscribe first.",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
		"state": sub.SubscribeState,
	})
}
