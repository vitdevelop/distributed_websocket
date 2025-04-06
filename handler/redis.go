package handler

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
)

var redisClient *redis.Client

func init() {
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	go handleRedisMessages()
}

func sendRedisMessage(message WsMessage) {
	jsonData, err := json.Marshal(InstanceMessage{
		InstanceName: instanceName,
		Message:      message,
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = redisClient.Publish(context.Background(), "messages", jsonData).Err()
	if err != nil {
		slog.Error(err.Error())
	}
}

func handleRedisMessages() {
	pubsub := redisClient.Subscribe(context.Background(), "messages")

	for msg := range pubsub.Channel() {
		message := InstanceMessage{}
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		if instanceName != message.InstanceName {
			broadcastUserMessage(User{}, message.Message)
		}
	}
}

func GetRedisConnectedUsers() []User {
	encodedUsers, err := redisClient.HGetAll(context.Background(), "users").Result()
	if err != nil {
		return []User{}
	}

	decodedUsers := make([]User, 0, len(encodedUsers))
	for _, encodedUser := range encodedUsers {
		user := User{}

		err = json.Unmarshal([]byte(encodedUser), &user)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		decodedUsers = append(decodedUsers, user)
	}
	return decodedUsers
}
