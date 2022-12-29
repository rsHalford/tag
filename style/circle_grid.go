package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func CircleGrid(c *generativeart.Canva) {
	c.SetLineWidth(4.0)
	c.Draw(arts.NewCircleGrid(4, 6))
}
