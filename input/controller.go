package input

type Controller interface {
	Move(dx, dy int)
	LeftClick()
	RightClick()
	Scroll(amount int)
	KeyDown(key string)
	KeyUp(key string)
}