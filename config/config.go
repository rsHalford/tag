/*
Package config creates the structure for the configuration of tag. This
includes setting the environment variables that the tag command and subcommands
use to operate according to the user's specification. These variables can also
be set using the config.toml file, that will be created automatically if it
does not already exist.

[general]
# location to save generated artwork (defaults to current working directory if unset)
directory = "" # $TAG_GENERAL_DIRECTORY

[canvas]
# width of the canvas to generate art with (defaults to "1920" pixels if unset)
width = "1920" # $TAG_CANVAS_WIDTH
# height of the canvas to generate art with (defaults to "1080" pixels if unset)
height = "1080" # $TAG_CANVAS_HEIGHT

[theme]
# default is the fallback theme to use for generative art (defaults to
# "rose_pine" if unset)
default = "rose_pine"
*/
package config

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type General struct {
	Directory string `toml:"directory" env:"DIRECTORY"`
}

type Canvas struct {
	Width  string `toml:"width" env:"WIDTH"`
	Height string `toml:"height" env:"HEIGHT"`
}

type Theme struct {
	Default string `toml:"default" env:"DEFAULT"`
}

// Config struct defines the config.toml and related environment variables.
type Config struct {
	General General `toml:"general" env-prefix:"TAG_GENERAL_"`
	Canvas  Canvas  `toml:"canvas" env-prefix:"TAG_CANVAS_"`
	Theme   Theme   `toml:"theme"  env-prefix:"TAG_THEME_"`
}

var cfg Config

// createCfgFile takes the configuration file path to determine if it exists,
// creating the file if missing.
func createCfgFile(cfgFile string) error {
	var f *os.File

	// If cfgFile represents a filepath that does not exist on the system,
	// one will be created.
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		f, err = os.Create(cfgFile)
		if err != nil {
			return fmt.Errorf("creating %v: %w", cfgFile, err)
		}

		defer f.Close()

		// The file is created with boilerplate for configuration options.
		configBoilerplate := fmt.Sprintln(`[general]
# location to save generated artwork (defaults to current working directory if unset)
directory = "" # $TAG_GENERAL_DIRECTORY

[canvas]
# width of the canvas to generate art with (defaults to "1920" pixels if unset)
width = "1920" # $TAG_CANVAS_WIDTH
# height of the canvas to generate art with (defaults to "1080" pixels if unset)
height = "1080" # $TAG_CANVAS_HEIGHT

[theme]
# default is the fallback theme to use for generative art (defaults to
# "rose_pine" if unset)
default = "rose_pine"`)

		_, err = f.WriteString(configBoilerplate)
		if err != nil {
			return fmt.Errorf("write to empty file: %w", err)
		}
	}

	return nil
}

// defineConfig assigns the file path for the configuration file, before
// checking the existence of the file and creating it if it doesn't exist.
func defineConfig() (cfgPath string, err error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("determining user configuration location: %v", err)
	}

	// Assign the config.toml filepath within the default root directory tag/,
	// to use in the user-specific configuration data.
	artCfgDir := cfgDir + "/tag"
	cfgPath = artCfgDir + "/config.toml"

	// Create the configuration directory if it doesn't already exist,
	// including a configuration file with acceptable options defined.
	if _, err = os.Stat(artCfgDir); os.IsNotExist(err) {
		var perm fs.FileMode = 0o755

		err = os.Mkdir(artCfgDir, perm)
		if err != nil {
			return "", fmt.Errorf("making new directory: %w", err)
		}

		err = createCfgFile(cfgPath)
		if err != nil {
			return "", fmt.Errorf("making new file: %w", err)
		}
	}

	// Create a configuration file if the directory is empty.
	err = createCfgFile(cfgPath)
	if err != nil {
		return "", fmt.Errorf("making new file: %w", err)
	}

	return cfgPath, nil
}

// Value takes a key and returns the matching value from the config.toml.
func Value(key string) string {
	cfgPath, err := defineConfig()
	if err != nil {
		fmt.Printf("finding configuration file: %v", err)
	}

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Printf("parsing configuration: %v", err)
	}

	switch key {
	case "general_directory":
		value := cfg.General.Directory

		return value

	case "canvas_width":
		value := cfg.Canvas.Width

		return value

	case "canvas_height":
		value := cfg.Canvas.Height

		return value

	case "theme_default":
		value := cfg.Theme.Default

		return value

	default:
		fmt.Println("Provided configuration key is not supported")
	}

	return ""
}
