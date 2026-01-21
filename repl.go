package main 

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


func cleanInput(text string) []string {
	//split user input into "words" based on whitespace 
	//trim any leading or trailing whitespave and lowercase the input

	lowered := strings.ToLower(text)
	
	//fields returns a slice of strings around each instance of one or more consecutive white spaces
	words := strings.Fields(lowered)
	return words
	
}

type cliCommand struct {
	name string 
	description string
	callback func()error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}