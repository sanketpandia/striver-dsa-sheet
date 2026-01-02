package datatypes

import "fmt"

func Run() {
	var a bool
	a = true

	var b string
	b = "Hello, World!"

	var c int32
	c = 24

	var d float32
	d = 3.279

	var e [5]int
	e[0] = 1
	e[1] = 2
	e[2] = 7
	e[3] = 4
	e[4] = 5

	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	var slices []string

	// slices[0] = "test"
	// slices[1] = "test2"

	slices = append(slices, "test3")

	fmt.Println("Slice Elements:")
	for _, e := range slices {
		fmt.Printf("%s\n", e)
	}

	fmt.Printf("\nData values: \nBoolean : %t\nString: %s\nInt: %d\nFloat: %f\n", a, b, c, d)

	for _, day := range days {
		fmt.Printf("\n%s", day)
	}

	fmt.Println("Numbers")
	for _, n := range e {
		fmt.Printf("\n%d", n)
	}

	fmt.Println("Days in reverse")
	for i := range days {
		fmt.Printf("\n%s", days[len(days)-1-i])
	}

}
