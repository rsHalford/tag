package style

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func DotsWave(c *generativeart.Canva) {
	c.Draw(arts.NewDotsWave(300))
}
