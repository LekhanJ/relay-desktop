package input

type Controller interface {
	Move(dx, dy int)

	LeftClick()
	RightClick()
	MiddleClick()

	Scroll(amount int)
	Zoom(delta float64)

	KeyDown(key string)
	KeyUp(key string)
}