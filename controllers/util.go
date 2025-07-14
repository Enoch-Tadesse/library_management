package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"os"
	"strconv"
	"strings"
)

// getNonEmptyInput takes in prompt and returns a non empty
// input string from the user. Incase of user wants to interrupt
// the process, it returns a false value.
func getNonEmptyInput(prompt string) (string, bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "m" {
			return "", false
		}
		if input == "" {
			fmt.Println("Input cannot be empty")
			continue
		}
		return input, true
	}
}

// getValidID take a function to call and delivers a
// better error handling that returns the validID with
// interruption signal
func getValidID(take func() (int, error)) (int, bool) {
	for {
		id, err := take()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if id == 0 {
			return 0, false
		}
		return id, true
	}
}

// membersPrinter takes a slice of Member to
// print in a formatted way
func membersPrinter(members []models.Member) {

	for _, member := range members {
		fmt.Println(strings.Repeat("-", 25))
		fmt.Printf("ID:        %-15d\n", member.ID)
		fmt.Printf("Name:      %-15s\n", member.Name)
	}
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Total:     %-15d\n\n", len(members))

}

// bookPrinter takes a slice of Book to
// print in a formatted way
func bookPrinter(books []models.Book) {
	for _, book := range books {
		fmt.Println(strings.Repeat("-", 25))
		fmt.Printf("ID:        %-15d\n", book.ID)
		fmt.Printf("Title:     %-15s\n", book.Title)
		fmt.Printf("Author:    %-15s\n", book.Author)
	}
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("Total:     %-15d\n\n", len(books))

}

// takeBookID prompts the user for a book id and
// returns a valid bookID and err
func takeBookID() (int, error) {
	input, ok := getNonEmptyInput("Book ID (or 'm' to return): ")
	if !ok {
		return 0, nil
	}
	bookID, err := strconv.Atoi(input)
	if err != nil || bookID <= 0 {
		return 0, fmt.Errorf("invalid book id: must be positive integer")
	}
	return bookID, nil
}

// takeMemberID prompts the user for a member id and
// returns a valid memberID and err
func takeMemberID() (int, error) {
	input, ok := getNonEmptyInput("Member ID (or 'm' to return): ")
	if !ok {
		return 0, nil
	}
	memberID, err := strconv.Atoi(input)
	if err != nil || memberID <= 0 {
		return 0, fmt.Errorf("invalid Member id: must be positive integer")
	}
	return memberID, nil
}
