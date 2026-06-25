package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LekhanJ/relay-desktop/input"
	"github.com/LekhanJ/relay-desktop/protocol"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},	
}

func main() {
	var controller input.Controller
	controller = &input.RobotGoController{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(controller, w, r)
	})

	log.Println("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(controller input.Controller, w http.ResponseWriter, r *http.Request) {
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

		var msg protocol.Message

		err = json.Unmarshal(message, &msg)
		if err != nil {
			continue
		}

		log.Printf("Type=%s dx=%f dy=%f", msg.Type, msg.DX, msg.DY)

		switch msg.Type {
			
		case protocol.MouseMove:
			controller.Move(msg.DX, msg.DY)

		case protocol.LeftClick:
			controller.LeftClick()

		case protocol.RightClick:
			controller.RightClick()	
		}
	}
}

func moveMouse(dx, dy int) {
	robotgo.MoveRelative(dx, dy)
}