package protocol

// MessageType represents the type of input event.
type MessageType string

const (
	MouseMove  MessageType = "mouse_move"
	MouseClick MessageType = "mouse_click"
	Scroll     MessageType = "scroll"
	Zoom       MessageType = "zoom"
	KeyDown    MessageType = "key_down"
	KeyUp      MessageType = "key_up"
)

// MouseButton represents a mouse button.
type MouseButton string

const (
	LeftButton   MouseButton = "left"
	RightButton  MouseButton = "right"
	MiddleButton MouseButton = "middle"
)

// Message represents a single input event sent over the WebSocket.
type Message struct {
	Type MessageType `json:"type"`

	// Mouse movement
	DX int `json:"dx,omitempty"`
	DY int `json:"dy,omitempty"`

	// Mouse click
	Button MouseButton `json:"button,omitempty"`

	// Mouse wheel scrolling.
	// Positive = up, Negative = down.
	Amount int `json:"amount,omitempty"`

	// Pinch zoom delta.
	// Positive = zoom in, Negative = zoom out.
	Delta float64 `json:"delta,omitempty"`

	// Keyboard key.
	Key string `json:"key,omitempty"`
}