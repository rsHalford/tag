/*
Lists all themes and styles available for artwork generation.

Usage:

	tag list [flags]

Aliases:

	list, ls, l

Flags:

	-h, --help           help for list
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all themes and styles",
	Long:    `Lists all themes and styles available for artwork generation.`,
	RunE:    listRun,
}

func listRun(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("please provide an argument or 'themes' of 'styles'.")
	} else {
		styles := `circle_grid
color_circle
color_circle_2
contour_line
dots_wave
maze
noise_line
random_shape`

		themes := `catppuccin_frappe
catppuccin_latte
catppuccin_macchiato
catppuccin_mocha
gruvbox_dark
gruvbox_dark_hard
gruvbox_dark_soft
gruvbox_light
gruvbox_light_hard
gruvbox_light_soft
one_dark
rose_pine
rose_pine_moon
rose_pine_dawn
solarized_dark
solarized_light
tokyo_night
tokyo_night_storm
tokyo_night_light`

		list := args[0]

		if list == "styles" {
			fmt.Println(styles)
			return nil
		} else if list == "themes" {
			fmt.Println(themes)
			return nil
		} else {
			return fmt.Errorf("please provide a valid argument for the list command")
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
