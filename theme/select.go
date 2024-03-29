package theme

import (
	"image/color"
)

func Select(theme string) (color.RGBA, []color.RGBA) {

	switch theme {
	case "catppuccin_frappe":
		base := CATPPUCCIN_FRAPPE_BASE
		schema := CATPPUCCIN_FRAPPE_COLORS

		return base, schema

	case "catppuccin_latte":
		base := CATPPUCCIN_LATTE_BASE
		schema := CATPPUCCIN_LATTE_COLORS

		return base, schema

	case "catppuccin_macchiato":
		base := CATPPUCCIN_MACCHIATO_BASE
		schema := CATPPUCCIN_MACCHIATO_COLORS

		return base, schema

	case "catppuccin_mocha":
		base := CATPPUCCIN_MOCHA_BASE
		schema := CATPPUCCIN_MOCHA_COLORS

		return base, schema

	case "gruvbox_dark":
		base := GRUVBOX_DARK_BASE
		schema := GRUVBOX_DARK_COLORS

		return base, schema

	case "gruvbox_dark_hard":
		base := GRUVBOX_DARK_HARD_BASE
		schema := GRUVBOX_DARK_COLORS

		return base, schema

	case "gruvbox_dark_soft":
		base := GRUVBOX_DARK_SOFT_BASE
		schema := GRUVBOX_DARK_COLORS

		return base, schema

	case "gruvbox_light":
		base := GRUVBOX_LIGHT_BASE
		schema := GRUVBOX_LIGHT_COLORS

		return base, schema

	case "gruvbox_light_hard":
		base := GRUVBOX_LIGHT_HARD_BASE
		schema := GRUVBOX_LIGHT_COLORS

		return base, schema

	case "gruvbox_light_soft":
		base := GRUVBOX_LIGHT_SOFT_BASE
		schema := GRUVBOX_LIGHT_COLORS

		return base, schema

	case "one_dark":
		base := ONE_DARK_BASE
		schema := ONE_DARK_COLORS

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

	case "solarized_dark":
		base := SOLARIZED_DARK_BASE
		schema := SOLARIZED_COLORS

		return base, schema

	case "solarized_light":
		base := SOLARIZED_LIGHT_BASE
		schema := SOLARIZED_COLORS

		return base, schema

	case "tokyo_night":
		base := TOKYO_NIGHT_BASE
		schema := TOKYO_NIGHT_COLORS

		return base, schema

	case "tokyo_night_storm":
		base := TOKYO_NIGHT_STORM_BASE
		schema := TOKYO_NIGHT_STORM_COLORS

		return base, schema

	case "tokyo_night_light":
		base := TOKYO_NIGHT_LIGHT_BASE
		schema := TOKYO_NIGHT_COLORS

		return base, schema

	default:
		base := ROSE_PINE_BASE
		schema := ROSE_PINE_COLORS

		return base, schema
	}
}
