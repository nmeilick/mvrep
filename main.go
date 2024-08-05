package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/kballard/go-shellquote"
	"github.com/urfave/cli/v2"
)

var version = ""

func main() {
	if version == "" {
		version = "0.0.0-dev"
	}

	app := &cli.App{
		Name:        "mvrep",
		Usage:       "Rename files and directories based on a pattern or fixed string substitution",
		Version:     version,
		HideVersion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"n"},
				Usage:   "Dry run: only show what would be changed",
			},
			&cli.BoolFlag{
				Name:    "fixed",
				Aliases: []string{"F"},
				Usage:   "Use fixed string substitution instead of regex",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "Interactive mode: ask for confirmation before renaming",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force overwrite of existing files",
			},
			&cli.BoolFlag{
				Name:    "shell",
				Aliases: []string{"s"},
				Usage:   "Output shell commands that would perform the renaming",
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"V"},
				Usage:   "Display version and exit",
			},
		},
		Action: renameFiles,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func renameFiles(c *cli.Context) error {
	if c.Bool("version") {
		fmt.Println("v" + version)
		return nil
	}

	args := c.Args().Slice()
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: mvrep [options] <pattern> <replacement> <files...>")
		return cli.Exit("", 1)
	}

	pattern := args[0]
	replacement := args[1]
	files := args[2:]

	if pattern == "" {
		fmt.Fprintln(os.Stderr, "Error: pattern cannot be empty")
		return cli.Exit("", 1)
	}

	var re *regexp.Regexp
	var err error
	if !c.Bool("fixed") {
		re, err = regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
			return cli.Exit("", 1)
		}
	}

	var useKeyboard bool
	if c.Bool("interactive") {
		if err := keyboard.Open(); err == nil {
			defer func() {
				_ = keyboard.Close()
			}()
			useKeyboard = true
		}
	}

	exitStatus := 0

	for _, file := range files {
		baseName := filepath.Base(file)
		dirName := filepath.Dir(file)

		var newName string
		if c.Bool("fixed") {
			newName = strings.ReplaceAll(baseName, pattern, replacement)
		} else {
			newName = re.ReplaceAllStringFunc(baseName, func(m string) string {
				matches := re.FindStringSubmatch(m)
				result := replacement
				for i, match := range matches {
					result = strings.ReplaceAll(result, fmt.Sprintf(`\%d`, i), match)
				}
				return result
			})
		}

		if newName == baseName {
			// Skip if no change
			continue
		}

		newPath := filepath.Join(dirName, newName)

		if c.Bool("dry-run") {
			fmt.Printf("Would rename: %s -> %s\n", file, newPath)
		} else if c.Bool("shell") {
			oldFileQuoted := shellquote.Join(file)
			newPathQuoted := shellquote.Join(newPath)
			if c.Bool("force") {
				fmt.Printf("mv -f %s %s\n", oldFileQuoted, newPathQuoted)
			} else {
				fmt.Printf("mv -n %s %s\n", oldFileQuoted, newPathQuoted)
			}
		} else {
			if c.Bool("interactive") {
				fmt.Printf("Rename %s to %s? (y/n): ", file, newPath)
				var response string
				if useKeyboard {
					for response == "" {
						char, key, err := keyboard.GetSingleKey()
						if err != nil {
							break
						}
						if key == keyboard.KeyEsc {
							fmt.Println("\nAborted!")
							os.Exit(0)
						}
						switch s := strings.ToLower(string(char)); s {
						case "y", "n":
							response = s
							fmt.Println()
						}
					}
				}
				if response == "" {
					fmt.Scanln(&response)
				}
				switch strings.ToLower(response) {
				case "y":
				default:
					fmt.Println("Skipped:", file)
					continue
				}
			}

			// Check if the new name already exists
			if _, err := os.Stat(newPath); !os.IsNotExist(err) {
				if c.Bool("force") {
					err := os.Remove(newPath)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error removing existing file %s: %v\n", newPath, err)
						exitStatus = 1
						continue
					}
				} else {
					fmt.Fprintf(os.Stderr, "Error: %s already exists\n", newPath)
					exitStatus = 1
					continue
				}
			}

			// Perform the actual renaming
			err := os.Rename(file, newPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error renaming %s: %v\n", file, err)
				exitStatus = 1
			} else {
				fmt.Printf("Renamed: %s -> %s\n", file, newPath)
			}
		}
	}

	return cli.Exit("", exitStatus)
}
