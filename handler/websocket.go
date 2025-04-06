package handler

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log/slog"
	"net"
	"net/http"
)

type Command uint

const (
	ConnectedUsers Command = 1
	Message                = 2
	CurrentUser            = 3
	Disconnect             = 4
)

type WsMessage struct {
	Command Command `json:"command,omitempty"`
	Data    any     `json:"data,omitempty"`
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	go handleWsConnection(conn)
}

type UserMessage struct {
	CurrentUser User   `json:"user,omitempty"`
	Message     string `json:"message,omitempty"`
}

func handleWsConnection(conn net.Conn) {
	currentUser := ConnectAvailableUser(conn)
	defer func() {
		conn.Close()
		ReturnAvailableUser(currentUser)
		broadcastUserMessage(currentUser, WsMessage{
			Command: Disconnect,
			Data:    currentUser,
		})
	}()

	sendWsMessage(conn, WsMessage{
		Command: CurrentUser,
		Data:    currentUser,
	})

	sendWsMessage(conn, WsMessage{
		Command: ConnectedUsers,
		Data:    GetConnectedUsers(),
	})
	broadcastUserMessage(currentUser, WsMessage{
		Command: ConnectedUsers,
		Data:    []User{currentUser},
	})

	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			return
		}

		if op == ws.OpText {
			message := UserMessage{}
			err = json.Unmarshal(msg, &message)
			if err != nil {
				slog.Error(err.Error())
				return
			}

			message.CurrentUser = currentUser
			broadcastUserMessage(currentUser, WsMessage{
				Command: Message,
				Data:    message,
			})
		}
	}

}

func sendWsMessage(conn net.Conn, message WsMessage) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = wsutil.WriteServerText(conn, jsonData)
	if err != nil {
		slog.Error(err.Error())
	}
}
