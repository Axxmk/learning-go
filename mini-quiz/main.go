package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the quiz!")

	//* Get user's name
	fmt.Print("Enter your name: ")
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Hello %s, nice to see you.\n", name)
	}

	//* Get user's age
	fmt.Print("Enter your age: ")
	var age uint
	if _, err := fmt.Scan(&age); err != nil {
		fmt.Println("Error:", err)
		return
	} else if age < 10 {
		fmt.Println("Sorry, you cannot play :(")
		return
	} else {
		fmt.Println("Let's start the quiz!")
	}

	//* Start the quiz
	score := 0
	numQuestions := 3
	scanner := bufio.NewScanner(os.Stdin)

	// q.1
	fmt.Println("1. In little red riding hood, who does the wolf dress up as?")
	scanner.Scan()
	if answer := scanner.Text(); answer == "Grandmother" || answer == "grandmother" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// q.2
	fmt.Println("2. How many colors are there in the rainbow?")
	scanner.Scan()
	if answer := scanner.Text(); answer == "Seven" || answer == "seven" || answer == "7" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	// q.2
	fmt.Println("3. Is Golang a type of OS? (Yes/No)")
	scanner.Scan()
	if answer := scanner.Text(); answer == "No" || answer == "no" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	//* Show result
	fmt.Printf("You scored %d out of %d.\n", score, numQuestions)
	percent := (float64(score) / float64(numQuestions)) * 100
	fmt.Printf("You got %.2f%%\n", percent)
}
