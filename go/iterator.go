package main

import (
	"fmt"
	"time"
)

const NUMBER_OF_ITERATIONS = 1000000000
const WHEN_PRINT = NUMBER_OF_ITERATIONS - 2
const NUMBER_OF_REPETITIONS = 100

func main() {
	var loopExecutions [NUMBER_OF_REPETITIONS]time.Duration
	var functionExecutions [NUMBER_OF_REPETITIONS]time.Duration

	for i := 0; i < NUMBER_OF_REPETITIONS; i++ {
		start := time.Now()

		for j := 0; j < NUMBER_OF_ITERATIONS; j++ {
			if j == WHEN_PRINT {
				end := time.Now()
				diff := end.Sub(start)
				fmt.Println(diff)
				loopExecutions[i] = diff
			}
			functionExecutions[i] = innerLoop()
		}
	}
	fmt.Println(loopExecutions)
	fmt.Println(functionExecutions)
}

func innerLoop() time.Duration {
	start := time.Now()

	for j := 0; j < NUMBER_OF_ITERATIONS; j++ {
		if j == WHEN_PRINT {
			end := time.Now()
			diff := end.Sub(start)
			fmt.Println(diff)
			return diff
		}
	}

	return 0
}
