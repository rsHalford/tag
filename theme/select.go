package theme

import (
	"image/color"
)

func Select(theme string) (color.RGBA, []color.RGBA) {

	switch theme {
	case "gruvbox_dark":
		base := GRUVBOX_DARK_BASE
		schema := GRUVBOX_DARK_COLORS

		return base, schema

	case "gruvbox_light":
		base := GRUVBOX_LIGHT_BASE
		schema := GRUVBOX_LIGHT_COLORS

		return base, schema

	case "rose_pine":
		base := ROSE_PINE_BASE
		schema := ROSE_PINE_COLORS

		return base, schema

	case "rose_pine_moon":
		base := ROSE_PINE_MOON_BASE
		schema := ROSE_PINE_MOON_COLORS

		return base, schema

	case "rose_pine_dawn":
		base := ROSE_PINE_DAWN_BASE
		schema := ROSE_PINE_DAWN_COLORS

		return base, schema

	default:
		base := ROSE_PINE_BASE
		schema := ROSE_PINE_COLORS

		return base, schema
	}
}
