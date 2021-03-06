package main

import "fmt"

type vertex struct {
	lat, lon float64
}

func main() {
	m := map[string]vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}
