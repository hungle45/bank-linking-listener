package main

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/pkg/kafka"
	"demo/bank-linking-listener/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

type BankLinkRequest struct {
	UserID   uint   `json:"user_id"`
	BankCode string `json:"bank_code"`
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func linkBankHandler(producer kafka.Producer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req BankLinkRequest
		err := ctx.BindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ResponseWithMessage(
				utils.ResponseStatusFail, err.Error()))
			return
		}

		jsonReq, err := json.Marshal(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ResponseWithMessage(
				utils.ResponseStatusFail, err.Error()))
			return
		}

		for _, topic := range producer.Topics() {
			message := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(jsonReq),
			}

			_, _, err = producer.SendMessage(message)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.ResponseWithMessage(
					utils.ResponseStatusFail, err.Error()))
				return
			}
		}

		ctx.JSON(http.StatusOK, utils.ResponseWithMessage(
			utils.ResponseStatusSuccess, "Bank linking request sent successfully"))
	}
}

func main() {
	config.LoadEnv("./.env")
	cfg := config.LoadConfig("./config.yml")
	fmt.Println(cfg)

	kafkaClient := kafka.NewClient(cfg)
	producer, err := kafka.NewProducer(kafkaClient, []string{"bank-linking-log"})
	if err != nil {
		log.Fatal("Failed to create kafka producer: ", err)
	}
	defer producer.Close()

	r := gin.Default()
	r.GET("/ping", pong)
	r.POST("/link-bank", linkBankHandler(producer))

	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", "localhost", "8081"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server running at %v:%v", "localhost", "8081")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
