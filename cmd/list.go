/*
Lists all themes available for artwork generation.

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
	Short:   "List all themes",
	Long:    `Lists all themes available for artwork generation.`,
	RunE:    listRun,
}

func listRun(cmd *cobra.Command, args []string) error {

	themes := `catppuccin_frappe
catppuccin_latte
gruvbox_dark
gruvbox_dark_hard
gruvbox_dark_soft
gruvbox_light
gruvbox_light_hard
gruvbox_light_soft
rose_pine
rose_pine_moon
rose_pine_dawn`

	fmt.Println(themes)

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
