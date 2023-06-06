package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// region helper functions
/* func loadScreen() {
	time.Sleep(time.Millisecond * 1000)
	println("#")
	time.Sleep(time.Millisecond * 1000)
	println("##")
	time.Sleep(time.Millisecond * 1000)
	println("###")
} */


func printWhiteSpace(ws int) {
	for i := 0; i < ws; i++ {
		println(" ")

	}
}

func initScreen() {
	fmt.Println(" +---+")
	fmt.Println(" |   |")
	fmt.Println(" O   |")
	fmt.Println("/|\\  | Lets play some Hangman!")
	fmt.Println("/ \\  |")
	fmt.Println("     |")
	fmt.Println("=========")
}

func firstFail() {
	fmt.Println("  Two  ")
	fmt.Println("Attempts")
	fmt.Println(" left ")
	fmt.Println("=========")
}
func secondFail() {
	fmt.Println(" +---+")
	fmt.Println("     |  One")
	fmt.Println("     | Attempt")
	fmt.Println("     | left")
	fmt.Println("     |")
	fmt.Println("     |")
	fmt.Println("=========")
}
func thirdFail() {
	fmt.Println(" +---+")
	fmt.Println(" |   |")
	fmt.Println(" O   |")
	fmt.Println("/|\\  | You are dead!")
	fmt.Println("/ \\  |")
	fmt.Println("     |")
	fmt.Println("=========")
}

func randNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

//endregion

// region game functions
func setName() string {
	input_name := ""
	initScreen()
	fmt.Println("Please enter your name:")
	fmt.Scanf("%s", &input_name)
	return input_name
}

func setDifficulty() int {
	difficulty := 0
	fmt.Println("Please enter a difficulty of 4, 6 or 8")
	fmt.Scanf("%d", &difficulty)
	for {
		if difficulty == 4 || difficulty == 6 || difficulty == 8 {
			break
		} else {
			fmt.Println("Only 4,6 or 8 allowed, please try again.")
			fmt.Scanf("%d", &difficulty)
		}
	}
	return difficulty
}

func guessing() string {
	guess := ""
	fmt.Printf("Please enter your guess\n")
	fmt.Scanf("%s", &guess)

	for {
		matched, _ := regexp.MatchString("[0-9]", guess)

		if len(guess) > 1 {
			fmt.Println("Please enter only one character")
			fmt.Scanf("%s", &guess)
		} else if matched {
			fmt.Println("No numeric characters allowed, only a-z, A-Z")
			fmt.Scanf("%s", &guess)
		} else {
			break
		}
	}
	return guess
}

//endregion

func main() {

	four_letter_words := []string{"Area", "Baby", "Book", "Food", "Game", "Girl", "Fish", "Risk", "Show", "Town"}
	six_letter_words := []string{"Bottom", "Broken", "Chance", "Column", "Design", "Driver", "Height", "Module", "Screen", "Theory"}
	eight_letter_words := []string{"Accident", "Aircraft", "Concrete", "Computer", "Disaster", "Employee", "Festival", "Feedback", "Pipeline", "Reporter"}

	input_name := setName()
	difficulty := setDifficulty()
	word := ""
	counter := 0
	solution := make([]string, difficulty)

	printWhiteSpace(2)

	fmt.Printf("Hello %s, lets start with a %d letter word!\n", input_name, difficulty)

	if difficulty == 4 {
		word = four_letter_words[randNumber()]
	} else if difficulty == 6 {
		word = six_letter_words[randNumber()]
	} else if difficulty == 8 {
		word = eight_letter_words[randNumber()]
	}
	printWhiteSpace(2)

	for i := 0; i < len(word); i++ {
		solution[i] = "_"
	}
	fmt.Println(solution)
	fmt.Printf("\n")

	printWhiteSpace(2)

	word_slice := strings.Split(word, "")

	for {
		guessed_char := guessing()
		flag := false

		for i := 0; i < len(word_slice); i++ {
			if word_slice[i] == guessed_char && word_slice[i] != "_" {
				solution[i] = guessed_char
				flag = true
			}
		}
		convert_solution_to_string := strings.Join(solution, "")

		if convert_solution_to_string == word {
			printWhiteSpace(2)
			fmt.Println("You Win!!!")
			fmt.Printf("Your word is %s", word)
			fmt.Printf("\n")
			break
		}
		if !flag {
			counter++
		}
		fmt.Println(solution)

		if counter == 1 {
			printWhiteSpace(2)
			firstFail()
		}
		if counter == 2 {
			printWhiteSpace(2)
			secondFail()
		}
		if counter == 3 {
			printWhiteSpace(2)
			thirdFail()
			break
		}
	}
}
