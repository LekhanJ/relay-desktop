package protocol

type MessageType string

const (
	MouseMove MessageType = "mouse_move"
	LeftClick MessageType = "left_click"
	RightClick MessageType = "right_click"
	Scroll MessageType = "scroll"
	KeyDown MessageType = "key_down"
	KeyUp MessageType = "key_up"
)

type Message struct {
	Type MessageType `json:"type"`
	DX int `json:"dx,omitempty"`
	DY int `json:"dy,omitempty"`
	Button string `json:"button,omitempty"` 
	Key string `json:"key,omitempty"`
	Amount int `json:"amount,omitempty"`
}