package main

import (
	"bufio"
	"fmt"
	"library"
	"mp"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager

var id int = 1

func handleLibCommand(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			music, _ := lib.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Artist, music.Source, music.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			music := library.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]}
			lib.Add(&music)
		} else {
			fmt.Println("Usage:lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			i, _ := strconv.ParseInt(tokens[2], 10, 0)
			lib.Remove(int(i))
		} else {
			fmt.Println("Usage: lib remove <index>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePalyCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage:play <name>")
		return
	}

	music := lib.Find(tokens[1])
	mp.Play(music.Source, music.Type)
}

func main() {
	fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib  
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play <name> -- Play the specified music
	`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter Command->")
		rawline, _, _ := r.ReadLine()
		line := string(rawline)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommand(tokens)
		} else if tokens[0] == "play" {
			handlePalyCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
