package main

import (
	"fmt"
	"strings"
)

var choice string
var output []rune
var output1 []rune

func main() {
	fmt.Println("enter a word (name,text)")
	fmt.Scan(&choice)
	fmt.Println("you entered", choice)

	runes := []rune(choice)
	for i := len(runes) - 1; i >= 0; i-- {
		output = append(output, runes[i])
	}
	fmt.Println("your string in reverse", string(output)) // use string to specify the output characters

	outputAsString := string(output)
	upperCasedString := strings.ToUpper(outputAsString)
	fmt.Println("upper case ", upperCasedString)

	normalString := reverse(string(output))
	fmt.Println("temp", normalString)
	// shift the string aski contenst with a number(random)

	fmt.Println("next conversion begine")
	shiftedString := shiftandreverse(string(choice))
	fmt.Println("next conversion is", shiftedString)

}

func reverse(output string) string {
	var workString strings.Builder
	workString.Grow(len(output))
	for i := len(output) - 1; i >= 0; i-- {
		workString.WriteByte(output[i])
	}
	return workString.String()
}

func shiftandreverse(choice string) (result string) {
	// shift the aski characters
	for _, v := range choice {
		result = string(v) + result
	}
	return
}
