package main

import "fmt"

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

func calcTotal() int {
	number := 982378273819287399
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	total := squares + cubes

	fmt.Println("Total Calculation Done!")
	return total
}

func main() {

	fmt.Println("Final output", calcTotal())
}
