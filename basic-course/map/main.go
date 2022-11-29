package main

import "fmt"

const GREEN = "green"

func main() {
	colors := map[string]string{
		"red": "#ff000",
		GREEN: "#4bf745",
	}

	fmt.Println("\n----- Define map -----")
	fmt.Println("\nAll key must be the same type")
	fmt.Println(colors)
	fmt.Printf("red: %v", colors["red"])
	printMap(colors)
	fmt.Println("\n----- Define map -----")
	fmt.Println("\n----- Delete map -----")
	delete(colors, GREEN)
	fmt.Printf("delete color green: %v", colors)
	fmt.Println(colors)
	colors[GREEN] = "#4bf745"

	otherColors := make(map[string]string)
	otherColors["white"] = "#ffffff"
}

func printMap(colors map[string]string) {
	fmt.Println("\n--- Custom print ---")
	for key, value := range colors {
		fmt.Printf("key %v: value %v\n", key, value)
	}
	fmt.Println("\n--- Custom print ---")
}
