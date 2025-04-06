package handler

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func broadcastUserMessage(user User, message WsMessage) {
	for _, userSession := range users {
		if userSession.user.Id == user.Id {
			continue
		}

		conn := userSession.conn
		sendWsMessage(conn, message)
	}

	// outside
	if user.Id == 0 {
		return
	}

	sendInstanceMessage(message)
}

type InstanceMessage struct {
	InstanceName string    `json:"instanceName"`
	Message      WsMessage `json:"message"`
}

func HandleInstanceMessage(w http.ResponseWriter, r *http.Request) {
	message := InstanceMessage{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if instanceName != message.InstanceName {
		broadcastUserMessage(User{}, message.Message)
	}

	w.WriteHeader(204)
}

func sendInstanceMessage(message WsMessage) {
	jsonData, err := json.Marshal(InstanceMessage{
		InstanceName: instanceName,
		Message:      message,
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/distributed/messages", bytes.NewReader(jsonData))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	client := http.Client{
		Timeout: 4 * time.Second,
	}

	_, err = client.Do(req)
	if err != nil {
		slog.Error(err.Error())
	}
}
