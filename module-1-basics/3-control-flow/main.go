package controlflow

import (
	"fmt"
	"math/rand"
)

func Run() {
	// Practise if-else conditions
	age := 28
	if age > 18 {
		fmt.Println("User is adult")
	} else if age < 18 {
		fmt.Println("User ain't adult")
	} else {
		fmt.Println("Soon!")
	}

	// With immediate initialization
	if a := rand.Intn(20); a < 10 {
		fmt.Printf("\nRand < 10: %d", a)
	} else {
		fmt.Printf("\nRand > 10: %d", a)
	}

	// Switch condition + with a for loop

	fmt.Println("Random integer based switch with expression")
	for range 5 {
		val := rand.Intn(20)
		switch val {
		case 12, 1, 2:
			fmt.Println("Winter Months")
		case 3, 4, 5, 6:
			fmt.Println("Summer Months")
		case 7, 8, 9, 10, 11:
			fmt.Println("Rain or Autumn")
		default:
			fmt.Println("Error value")
		}
	}

	fmt.Println("\nSwitch without expression")
	for range 5 {
		val := rand.Intn(5)
		switch {
		case val%2 == 0:
			fmt.Println("Number is even")
		default:
			fmt.Printf("Number is not even and could be odd: %d\n", val)
		}
	}

	// Range experiment and for loops
	// Mocking a while loop
	i := 0
	for i < 5 {
		fmt.Printf("Iteration %d\n", i)
		i++
	}

	j := 0
	for {
		fmt.Printf("Looping through %d\n", j)
		if j > 5 {
			break
		}
		j++
	}

	fmt.Println("Writing the number guessing game: Between 1-20")

	var maxGuesses int
	fmt.Println("Enter maximum number of attempts:")
	fmt.Scanln(&maxGuesses)

	num := rand.Intn(20)

	counter := 1
outerloop:
	for {

		var x int
		fmt.Println("Enter your guess:")
		fmt.Scanln(&x)

		if x == num {
			fmt.Println("Congratulations! You guessed the number correctly")
			break outerloop
		} else if x < num {
			fmt.Println("Too cold")
		} else if x > num {
			fmt.Println("Too hot")
		}
		if counter >= maxGuesses {
			fmt.Println("Max number of guesses exhausted")
			fmt.Printf("Number was %d", num)
			break outerloop
		}
		counter++
	}

}
