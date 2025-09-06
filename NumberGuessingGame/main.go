package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	secret := rand.Intn(50) + 1
	var guess int
	guesses := make(chan int)
	timeout := time.After(15 * time.Second) 

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Guess a number between 1 and 50 (you have 15 seconds):")

	go func() {
		for {
			fmt.Print("Enter your guess: ")
			_, err := fmt.Scan(&guess)
			if err == nil {
				guesses <- guess
			}
		}
	}()

	for {
		select {
		case g := <-guesses:
			if g < secret {
				fmt.Println("Too low! Try again.")
			} else if g > secret {
				fmt.Println("Too high! Try again.")
			} else {
				fmt.Println("Correct! You win!")
				return
			}
		case <-timeout:
			fmt.Println("\nTime’s up! You lost. The number was:", secret)
			return
		}
	}
}
