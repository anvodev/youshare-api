package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"youshare-api.anvo.dev/internal/data"
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

var videoCreated = make(chan data.Video)

func (app *application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.logger.Print("error upgrading connection", err)
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	register <- ws
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			app.logger.Println("read fail:", err)
			break
		}
	}
	unregister <- ws
}

func (app *application) broadcastWs() {
	for {
		select {
		case client := <-register:
			clients[client] = true
			app.logger.Print("Client connected")
		case client := <-unregister:
			if _, ok := clients[client]; ok {
				delete(clients, client)
				client.Close()
				app.logger.Print("Client disconnected")
			}
		case msg := <-videoCreated:
			for client := range clients {
				err := client.WriteJSON(envelop{"video": msg})
				if err != nil {
					unregister <- client
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
