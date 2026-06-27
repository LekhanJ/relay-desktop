package input

import (
	"github.com/go-vgo/robotgo"
)

type RobotGoController struct {
	zoomAccumulator float64
}

func (c *RobotGoController) Move(dx, dy int) {
	x, y := robotgo.Location()
	robotgo.Move(x+dx, y+dy)
}

func (c *RobotGoController) LeftClick() {
	robotgo.Click("left")
}

func (c *RobotGoController) MiddleClick() {
	robotgo.Click("center")
}

func (c *RobotGoController) RightClick() {
	robotgo.Click("right")
}

func (c *RobotGoController) Scroll(amount int) {
	if amount > 0 {
        robotgo.ScrollDir(amount, "up")
    } else if amount < 0 {
        robotgo.ScrollDir(-amount, "down")
    }
}

func (c *RobotGoController) KeyDown(key string) {
	robotgo.KeyDown(key)
}

func (c *RobotGoController) KeyUp(key string) {
	robotgo.KeyUp(key)
}

func (c *RobotGoController) Zoom(delta float64) {
	c.zoomAccumulator += delta

	const threshold = 0.02

	robotgo.KeyToggle("ctrl", "down")
	defer robotgo.KeyToggle("ctrl", "up")

	for c.zoomAccumulator >= threshold {
		robotgo.ScrollDir(1, "up")
		c.zoomAccumulator -= threshold
	}

	for c.zoomAccumulator <= -threshold {
		robotgo.ScrollDir(1, "down")
		c.zoomAccumulator += threshold
	}
}