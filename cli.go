package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	var r, err = os.Open(DefaultConfigPath())
	check(err)

	stash, err := buildStashFromBuffer(r)
	check(err)

	args := os.Args[1:]
	switch len(args) {
	case 0:
		ListEntries(stash)
	case 1:
		GetEntry(stash, args)
	default:
		AddEntry(stash, args)
	}
}

//Add a new entry to the stash by taking the first argument as the key, and joining the rest
//as the value
func AddEntry(stash *Stash, args []string) {
	command := strings.Join(args[1:], " ")
	key := args[0]
	stash.Add(key, command)
	fmt.Printf("Added: %v => %v", key, command)

	f, err := os.Create(DefaultConfigPath())
	check(err)
	defer f.Close()
	stash.SaveStashToWriter(f)
	f.Sync()
}

//List all entries in the stash
func ListEntries(stash *Stash) {
	fmt.Print(stash.Format())
}

//Get an entry out of the stash and copy its value to the clipboard
func GetEntry(stash *Stash, args []string) {
	command := stash.Get(args[0])
	clipboard.WriteAll(command)
	fmt.Printf("Copied: %v\n", command)
}

func check(e error) {
	if e != nil {
		fmt.Print("error occured", e)
		os.Exit(1)
	}
}
