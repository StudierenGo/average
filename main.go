// Calculate the average of a list of numbers read from a file.
package main

import "fmt"

func main() {
	var sum float64
	numbers := [3]float64{10.5, 20.3, 30.2}

	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Average: %.2f\n", sum/float64(len(numbers)))
}
