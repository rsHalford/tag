package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func RandomShape(c *generativeart.Canva) {
	c.Draw(arts.NewRandomShape(111))
}
