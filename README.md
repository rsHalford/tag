<p align="center">
  <a href="https://pkg.go.dev/github.com/rsHalford/tag">
    <img src="https://pkg.go.dev/badge/github.com/rsHalford/tag.svg" alt="Go Reference">
  </a>
  <a href="https://goreportcard.com/report/github.com/rsHalford/tag">
    <img src="https://goreportcard.com/badge/github.com/rsHalford/tag" alt="Go Report Card">
  </a>
</p>

---

# About

tag generates artwork from your favourite terminal themes.

```
Usage:
  tag [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Create a new generative artwork
  help        Help about any command

Flags:
  -h, --help      help for tag
  -v, --version   version for tag

Use "tag [command] --help" for more information about a command.
```

Go get things done and checked off the list.

---

# Getting Started

## Requirements

All you need is to have Go [installed](https://go.dev/dl/) to compile tag.

## Installation

To install tag, all you have to do is run the `go install` command.

```sh
$ go install github.com/rsHalford/tag@latest
```

### Flake

If you're into Nix, tag has been setup so that you can just add the repo's URL to your flake.nix inputs. Then the overlay can be called by `environment.systemPackages` or per user with `home.packages`.

```nix
inputs = {
  nixpkgs.url = "nixpkgs/nixos-unstable";

  tag-flake = {
    url = "github:rsHalford/tag";
    inputs.nixpkgs.follows = "nixpkgs";
  };
};
```

Instructions on how to add tag as a flake can vary wildly depending on the system configuration. You can view [my dotfiles](https://github.com/rsHalford/dotfiles) to see a working example, and message me or create an issue on either repository if you need additional help getting setup.

## Configuration

You can edit the `config.toml` to set your preferred default settings, helping shorten your most used `tag` arguments.

On Linux this file will be read from `${XDG_CONFIG_HOME:-$HOME/.config}/tag/config.toml`.

```toml
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
```

---

# Contributing

To help contribute to tag, you could either send in your feature requests as an issue or take it upon yourself to send in a pull request after following the [CONTRIBUTING](https://github.com/rsHalford/tag/blob/main/CONTRIBUTING.md) guide.

Thanks in advance for taking an interest!

## Development Environment

tag development is setup to utilise the following tools;

- [git-chglog](https://github.com/git-chglog/git-chglog)
  - For updating the CHANGELOG.md
- [gopls](https://github.com/golang/tools/blob/master/gopls/README.md)
  - Go language server

### [direnv](https://direnv.net/)

Combined with nix flakes and [nix-direnv](https://github.com/nix-community/nix-direnv), it is possible to create a `.envrc` file to make a portable isolated environment for development.

Meaning the above tools can be installed and setup just for this project by using the included `.envrc` found in the root of the project directory.

```sh
use flake
layout go
```

---

# Acknowledgment

tag wouldn't be possible without the below libraries.

- [generativeart](https://github.com/jdxyw/generativeart)
- [cobra](https://github.com/spf13/cobra)

---

# Licence

tag is released under the MIT License.

See [LICENSE](https://github.com/rsHalford/tag/blob/main/LICENSE).
