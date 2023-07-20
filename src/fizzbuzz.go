package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) {
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}

// Q. この関数は何をしているのか？のように関数の説明をCopilotに聞いてみましょう。
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number: ")

	if scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid integer.")
			os.Exit(1)
		}

		fizzBuzz(num)
	} else {
		fmt.Println("Failed to read input.")
		os.Exit(1)
	}
}
