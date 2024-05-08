package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p:=price(); {
		case p < 2:
			fmt.Println("cheatp item")
		case p < 10:
			fmt.Println("medium priced item")
		default:
			fmt.Println("expensive item")
	}

	ticket := Economy
	switch ticket {
		case Economy:
			fmt.Println("Economy seating")
		case Business:
			fmt.Println("Business seating")
		case FirstClass:
			fmt.Println("Firstclass seating")
		default:
			fmt.Println("Unknown seating")
	}

}
