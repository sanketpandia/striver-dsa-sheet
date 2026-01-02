package helloworld

import "fmt"

func Run() {
	var name string

	fmt.Println("Enter Name:")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s \n", name)

	var year int
	fmt.Println("Enter year and name")
	fmt.Scanf("%d %s", &year, &name)
	fmt.Printf("\n Year: %d, Name: %s", year, name)
}
