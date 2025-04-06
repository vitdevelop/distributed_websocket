package main

import (
	"context"
	"distributedwebsocket/handler"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"net/http"
	"syscall"
)

func main() {
	ln := listenerConfig()

	http.HandleFunc("/", handler.HandleDefaultPage)
	http.Handle("/{filename}", http.FileServer(http.Dir("./www")))
	http.Handle("/heroes/{filename}", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/ws", handler.HandleWs)
	http.HandleFunc("POST /distributed/messages", handler.HandleInstanceMessage)
	http.HandleFunc("GET /distributed/users", handler.HandleInstanceConnectedUsers)

	log.Fatal(http.Serve(ln, nil))
}

func listenerConfig() net.Listener {
	lc := net.ListenConfig{
		Control: func(network, address string, conn syscall.RawConn) error {
			var operr error
			if err := conn.Control(func(fd uintptr) {
				operr = syscall.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			}); err != nil {
				return err
			}
			return operr
		},
	}

	ln, err := lc.Listen(context.Background(), "tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	return ln
}
