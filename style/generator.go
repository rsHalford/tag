package style

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/rsHalford/tag/config"
	"github.com/rsHalford/tag/theme"
)

func Generator(wg *sync.WaitGroup, s, th string) {
	defer wg.Done()

	width, _ := strconv.Atoi(config.Value("canvas_width"))
	height, _ := strconv.Atoi(config.Value("canvas_height"))

	c := generativeart.NewCanva(width, height)

	base, schema := theme.Select(th)

	c.SetBackground(base)
	c.FillBackground()
	c.SetColorSchema(schema)

	now := time.Now()
	unix := now.Unix()
	time := strconv.FormatInt(unix, 10)

	switch s {
	case "color_circle":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		ColorCircle(c)

	case "color_circle_2":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		ColorCircle2(c)

	case "contour_line":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		ContourLine(c)

	case "maze":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		Maze(c, schema)

	case "noise_line":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		NoiseLine(c)

	case "random_shape":
		fmt.Printf("Generating %s_%s-%s.png...\n", th, s, time)
		RandomShape(c)

	default:
		fmt.Printf("%s is not a supported style\n", s)
	}

	if config.Value("general_directory") != "" {
		saveDir := strings.TrimSuffix(config.Value("general_directory"), "/")
		saveLocation := saveDir + "/" + th + "_" + s + "-" + time + ".png"

		c.ToPNG(saveLocation)

	} else {
		c.ToPNG(th + "_" + s + "-" + time + ".png")
	}
}
