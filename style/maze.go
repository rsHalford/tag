package style

import (
	"image/color"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func Maze(c *generativeart.Canva, schema []color.RGBA) {
	c.SetLineWidth(1)
	for i := range schema {
		c.SetLineWidth(float64((i+1)/3 + 1))
		c.Draw(arts.NewMaze(i*10 + 100))
		c.SetLineColor(schema[i])
	}
}
