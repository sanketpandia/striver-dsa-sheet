package main

import (
	mapper "./gocsvfilemapper"
	"fmt"
)

func main() {
	path := "./sample_csv.csv"
	csv, err := mapper.NewCSVFilter(path)

	if err != nil {
		fmt.Println("Failed to parse file")
		fmt.Println(err)
	}

	fmt.Printf("Length: %d", csv.Len)
}
