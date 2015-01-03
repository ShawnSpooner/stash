package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func main() {
	var r, err = os.Open(DefaultConfigPath())
	check(err)

	stash, err := buildStashFromBuffer(r)
	check(err)

	app := cli.NewApp()
	app.Name = "stash"
	app.Usage = "stash away usefull commands"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		switch len(c.Args()) {
		case 0:
			ListEntries(stash, c)
		case 1:
			GetEntry(stash, c)
		default:
			AddEntry(stash, c)
		}
	}
	app.Run(os.Args)
}

//Add a new entry to the stash by taking the first argument as the key, and joining the rest
//as the value
func AddEntry(stash *Stash, c *cli.Context) {
	command := strings.Join(c.Args()[1:], " ")
	key := c.Args()[0]
	stash.Add(key, command)
	fmt.Printf("Added: %v => %v", key, command)

	f, err := os.Create(DefaultConfigPath())
	check(err)
	defer f.Close()
	stash.SaveStashToWriter(f)
	f.Sync()
}

//List all entries in the stash
func ListEntries(stash *Stash, c *cli.Context) {
	fmt.Print(stash.Format())
}

//Get an entry out of the stash and copy its value to the clipboard
func GetEntry(stash *Stash, c *cli.Context) {
	command := stash.Get(c.Args()[0])
	clipboard.WriteAll(command)
	fmt.Printf("Copied: %v", command)
}

func check(e error) {
	if e != nil {
		fmt.Print("error occured", e)
		os.Exit(1)
	}
}
