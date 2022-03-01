package main

import "fmt"

func main() {

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"white": "#ffffff",
	// }

	// var colors map[string]string

	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["white"] = "#fffff"

	printMap(colors)

	delete(colors, "white")

	fmt.Println(colors)

}

// looping over a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
