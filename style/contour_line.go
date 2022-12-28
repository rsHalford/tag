package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func ContourLine(c *generativeart.Canva) {
	c.Draw(arts.NewContourLine(999))
}
