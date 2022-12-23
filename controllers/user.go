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

type SendNotificationRequest struct {
	EveryMinute int
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
		subscription.CourseID = ""
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
	// Todo: Need add authorization
	vapidPublicKey := os.Getenv("VAPID_PUBLIC_KEY")
	vapidPrivateKey := os.Getenv("VAPID_PRIVATE_KEY")

	var req SendNotificationRequest
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
		return
	}

	// get all subscription mapping with every minute
	var subs []models.Subscription
	tx := db.Orm.Find(&subs, "every_minute = ? and subscribe_state = ?", req.EveryMinute, true)
	if tx.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": fmt.Sprintf("error when get subs with every minute %d", req.EveryMinute),
		})
	}

	if tx.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "No sub to run",
		})
		return
	}

	var resp *http.Response
	for _, v := range subs {
		card, err := getRandomCardFromCourse(v.CourseID, v.UserID)
		if err != nil {
			fmt.Println("error when random card with course id " + v.CourseID + ", user id " + v.UserID)
		} else {
			// Decode subscription
			s := &webpush.Subscription{}
			json.Unmarshal([]byte(v.Sub), s)

			// Send Notification
			resp, err = webpush.SendNotification([]byte(card.Terminology+" / "+card.Definition), s, &webpush.Options{
				Subscriber:      "example@example.com",
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
		}
	}
	defer resp.Body.Close()

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func getRandomCardFromCourse(courseId string, userId string) (*models.Card, error) {
	var card models.Card
	if courseId != "" {
		tx := db.Orm.Raw("select * from cards where course_id = ? order by random() limit 1", courseId).Scan(&card)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := db.Orm.Raw("select * from cards inner join courses on cards.course_id = courses.id where courses.user_id = ? order by random() limit 1", userId).Scan(&card)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &card, nil
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
			"error":  tx.Error.Error(),
		})
		return
	}

	if tx.RowsAffected > 0 {
		sub.SubscribeState = req.State
		db.Orm.Save(&sub)
	} else {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "cannot get subscription, please subscribe first.",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
		"state":  sub.SubscribeState,
	})
}
