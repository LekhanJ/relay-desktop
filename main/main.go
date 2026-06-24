package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
	DX float64 `json:"dx"`
	DY float64 `json:"dy"`	
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},	
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	log.Println("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	log.Println("Android Connected")

	for {
		_, message, err := conn.ReadMessage();
		if err != nil {
			log.Println(err)
			break
		}

		var msg Message

		err = json.Unmarshal(message, &msg)
		if err != nil {
			continue
		}

		log.Printf("Type=%s dx=%f dy=%f", msg.Type, msg.DX, msg.DY)

		if msg.Type == "move" {
			moveMouse(int(msg.DX), int(msg.DY))
		}
	}
}

func moveMouse(dx, dy int) {
	robotgo.MoveRelative(dx, dy)
}