package main 

import (
	"strings"
)

func cleanInput(text string) []string {
	//split user input into "words" based on whitespace 
	//trim any leading or trailing whitespave and lowercase the input

	lowered := strings.ToLower(text)
	
	//fields returns a slice of strings around each instance of one or more consecutive white spaces
	words := strings.Fields(lowered)
	return words
	
}