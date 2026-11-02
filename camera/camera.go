package camera

import (
	"math"
)

type Camera struct {
	X, Y, targetX, targetY float64
}

func NewCamera(x, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
		targetX: 0,
		targetY: 0,
	}
}

func (c *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	x, y := -targetX + screenWidth/2.0, -targetY + screenHeight/2.0

	// if c.targetX - c.X < 0.1 && c.targetY - c.Y < 0.1 {
		c.targetX = x
		c.targetY = y
	// }
}

func (c *Camera) Update() {
	if math.Abs(c.targetX - c.X) >2 || math.Abs(c.targetY - c.Y) > 2 {
		if c.X < c.targetX {
			if  math.Abs(c.X - c.targetX) > 1 {
				c.X += 1
			}
		} else {
			if  math.Abs(c.targetX - c.X) > 1 {
				c.X -= 1
			}
		}
		if c.Y < c.targetY {
			if  math.Abs(c.Y - c.targetY) > 1 {
				c.Y += 1
			}
		} else {
			if  math.Abs(c.targetY - c.Y) > 1 {
				c.Y -= 1
			}
		}
	}

	
}

func (c *Camera) Constrain(tilemapWidthPixels, tilemapHeightPixels, screenWidth, screenHeight float64) {
	c.X = math.Min(c.X, 50.0)
	c.Y = math.Min(c.Y, 50.0)

	c.X = math.Max(c.X, screenWidth-tilemapWidthPixels - 50)
	c.Y = math.Max(c.Y, screenHeight-tilemapHeightPixels - 50)
}