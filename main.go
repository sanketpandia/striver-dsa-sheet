package main

import (
	"fmt"
	"strings"

	csvmapper "sanketpandia/striver-dsa-sheet/go-csv-file-mapper/gocsvfilemapper"
	helloworld "sanketpandia/striver-dsa-sheet/module-1-basics/1-hello-world"
	datatypes "sanketpandia/striver-dsa-sheet/module-1-basics/2-data-types"
	controlflow "sanketpandia/striver-dsa-sheet/module-1-basics/3-control-flow"
	functions "sanketpandia/striver-dsa-sheet/module-1-basics/4-functions-memory"
	patterns "sanketpandia/striver-dsa-sheet/module-1-basics/5-pattern-problems"
)

var programs = []struct {
	name string
	run  func()
}{
	{"Hello World", helloworld.Run},
	{"Data Types", datatypes.Run},
	{"Control Flow", controlflow.Run},
	{"Functions & Memory", functions.Run},
	{"Pattern Problems", patterns.Run},
	{"CSV File Mapper", runCSVMapper},
}

func runCSVMapper() {
	path := "./go-csv-file-mapper/sample_csv.csv"
	csv, err := csvmapper.NewCSVFilter(path)
	if err != nil {
		fmt.Println("Failed to parse file:")
		fmt.Println(err)
		return
	}

	fmt.Println("Headers for Interactive filtering:")
	fmt.Println(strings.Join(csv.Headers, ", "))

	for {

		fmt.Println("Enter the column index (-1 to exit)")
		var i int
		var text string
		fmt.Scanln(&i)

		if i == -1 {
			break
		}
		fmt.Println("Enter text to filter with:")
		fmt.Scanln(&text)

		csv.Filter(i, text)
	}

	fmt.Printf("Loaded CSV with %d records\n", csv.Len)
}

func main() {
	for {
		fmt.Println("\n=== Striver DSA Sheet ===")
		for i, p := range programs {
			fmt.Printf("%d. %s\n", i+1, p.name)
		}
		fmt.Println("0. Exit")
		fmt.Print("\nEnter choice: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			fmt.Println("Goodbye!")
			break
		}

		if choice < 1 || choice > len(programs) {
			fmt.Println("Invalid choice, try again.")
			continue
		}

		fmt.Printf("\n--- Running: %s ---\n\n", programs[choice-1].name)
		programs[choice-1].run()
	}
}
