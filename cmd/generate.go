/*
Generates artwork using the provided art styles and theme palette.

Usage:

	tag generate [flags]

Aliases:

	generate, gen, g

Flags:

	-h, --help           help for generate
	-t, --theme string   name of the artwork's theme
*/
package cmd

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/rsHalford/tag/config"
	"github.com/rsHalford/tag/style"
	"github.com/spf13/cobra"
)

var (
	themeOpt string
)

// generateCmd represents the edit command.
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "g"},
	Short:   "Create a new generative artwork",
	Long:    `Generates artwork using the provided art styles and theme palette.`,
	RunE:    generateRun,
}

func generateRun(cmd *cobra.Command, args []string) error {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup
	var th string

	if themeOpt != "" {
		th = themeOpt
	} else if config.Value("theme_default") != "" {
		th = config.Value("theme_default")
	} else {
		th = "rose_pine"
	}

	for _, s := range args {
		wg.Add(1)
		go style.Generator(&wg, s, th)
	}

	wg.Wait()

	fmt.Println("Artwork generation finished!")

	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// The --theme flag argument determines what theme palette to use with generation.
	generateCmd.Flags().StringVarP(&themeOpt, "theme", "t", "", "name of the artwork's theme")
}
