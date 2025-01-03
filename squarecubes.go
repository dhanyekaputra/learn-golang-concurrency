package main

import "fmt"

// function takes int as first parameter
// the second param is channel, we'll use that later for storing our calc value
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	fmt.Println("Square Calculation Done!")
	squareop <- sum

}

// function takes int as first parameter
// the second param is channel, we'll use that later for storing our calc value
func calcCubes(number int, cubeop chan int) {
	sum := 0

	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}

	fmt.Println("Cube Calculation Done!")
	cubeop <- sum

}

// goroutine and channel implementation
func calcTotal() int {
	number := 982378273819287399

	// sqrch and cubech channel created
	sqrch := make(chan int)
	cubech := make(chan int)

	// calcSquares goroutine will start
	go calcSquares(number, sqrch)

	// calcCubes goroutine will start
	go calcCubes(number, cubech)

	// program will stop until
	// sqrch channel and cubech channel
	// receiving data from both goroutine
	squares, cubes := <-sqrch, <-cubech
	total := squares + cubes

	fmt.Println("Total Calculation Done!")
	return total
}
