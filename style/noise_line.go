package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func NoiseLine(c *generativeart.Canva) {
	c.Draw(arts.NewNoiseLine(3456))
}
