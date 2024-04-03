package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	register   = make(chan *websocket.Conn)
	unregister = make(chan *websocket.Conn)
	clients    = make(map[*websocket.Conn]bool)
)

var videoCreated = make(chan string)

func (app *application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.logger.Print("error upgrading connection", err)
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	register <- ws
	app.logger.Print("Client connected")
	defer func() {
		unregister <- ws
		ws.Close()
		app.logger.Print("Client disconnected")
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			app.logger.Println("read fail:", err)
			break
		}
	}
}

func (app *application) broadcastWs() {
	for {
		select {
		case client := <-register:
			clients[client] = true
		case client := <-unregister:
			if _, ok := clients[client]; ok {
				delete(clients, client)
				client.Close()
			}
		case msg := <-videoCreated:
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					unregister <- client
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
