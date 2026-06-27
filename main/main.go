package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LekhanJ/relay-desktop/input"
	"github.com/LekhanJ/relay-desktop/protocol"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	var controller input.Controller = &input.RobotGoController{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(controller, w, r)
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(controller input.Controller, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Android connected")

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var msg protocol.Message
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Println("Invalid message:", err)
			continue
		}

		log.Printf("Received: %+v\n", msg)

		switch msg.Type {

		case protocol.MouseMove:
			controller.Move(msg.DX, msg.DY)

		case protocol.MouseClick:
			switch msg.Button {
			case protocol.LeftButton:
				controller.LeftClick()

			case protocol.RightButton:
				controller.RightClick()

			case protocol.MiddleButton:
				controller.MiddleClick()
			}

		case protocol.Scroll:
			controller.Scroll(msg.Amount)

		case protocol.Zoom:
			controller.Zoom(msg.Delta)

		case protocol.KeyDown:
			controller.KeyDown(msg.Key)

		case protocol.KeyUp:
			controller.KeyUp(msg.Key)

		default:
			log.Printf("Unknown message type: %s\n", msg.Type)
		}
	}
}