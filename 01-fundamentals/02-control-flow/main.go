package main

import "fmt"

func main() {
	// TODO: classic three-clause for loop, i from 1 to 20 inclusive.
	// Inside, use a tagless switch to print FizzBuzz output (see README).
	for i := 1; i <= 20; i++ {

		switch {
		case i%15 == 0 && i%3 == 0:
			fmt.Printf("FizzBuzz\n")
		case i%3 == 0:
			fmt.Printf("Fizz\n")
		case i%5 == 0:
			fmt.Printf("Buzz\n")
		default:
			fmt.Printf("%d\n", i)
		}

	}

	// TODO: condition-only for loop ("while" form) counting down from 5
	// to 1, printing "T-minus N" each time, then "Liftoff!" at the end.
	var j int = 5
	for j >= 0 {
		fmt.Printf("%d\n", j)
		j--
	}

	// TODO: nested for loops (i, j each 1..10) to find the first pair
	// whose product is 40. Use a labeled break to exit both loops as
	// soon as it's found, then print "Found pair: %d x %d = %d".
	for i := range 11 {
		for j := range 11 {
			if i*j == 40 {
				fmt.Printf("Found pair: %d x %d = %d\n", i, j, 40)
			}
		}
	}

	// TODO: using an if with a short init statement, check whether the
	// sum of the pair found above is even or odd, and print
	// "Sum %d is even" or the odd equivalent.
	if isEven := 13%2 == 0; isEven {
		fmt.Printf("Sum %d is even\n", 13)
	} else {
		fmt.Printf("Sum %d is odd\n", 13)
	}

}
