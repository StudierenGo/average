// Calculate the average of a list of numbers read from a file.
package main

import (
	"average/datafile"
	"fmt"
	"log"
)

func main() {
	var sum float64
	numbers, err := datafile.GetFloats("data/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Average: %.2f\n", sum/float64(len(numbers)))

	votes, err := datafile.GetStrings("data/votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(votes)
}
