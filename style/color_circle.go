package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func ColorCircle(c *generativeart.Canva) {
	c.Draw(arts.NewColorCircle(1234))
}
