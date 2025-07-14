# Library Management System Documentation

## Overview

This project implements a simple library management system with a command-line interface. It allows users to interact with a library's collection of books and members through various operations, such as adding or removing books and members, borrowing and returning books, and listing available or borrowed books. The system uses a `LibraryManager` interface (defined in `library_services.go`) to manage the underlying data.

Key features include:
- Adding new books and members
- Removing existing books and members
- Borrowing and returning books
- Listing all members, available books, and books borrowed by a specific member

## Main Components

### Menu System
The `Menu` map in `library_controller.go` defines the available user options and associates each option with a corresponding function. The `DisplayMenu` function prints these options in a formatted menu for the user to choose from.

### User Input Handling
Utility functions in `util.go` are responsible for handling and validating user input:
- `getNonEmptyInput`: Ensures the user provides a non-empty string input
- `getValidID`: Ensures the user provides a valid positive integer ID

### Library Operations
Functions in `library_controller.go` handle the core operations of the library system by interacting with the `LibraryManager` interface:
- `AddBook`: Adds a new book to the library
- `AddMember`: Adds a new member to the library
- `RemoveBook`: Removes a book from the library
- `RemoveMember`: Removes a member from the library
- `BorrowBook`: Allows a member to borrow a book
- `ReturnBook`: Allows a member to return a borrowed book
- `ListMembers`: Lists all registered members
- `ListAllAvailableBooks`: Lists all books that are currently available
- `ListAllBorrowedByAMember`: Lists all books borrowed by a specific member

### Data Display
Helper functions in `util.go` are used to format and print lists of books and members:
- `bookPrinter`: Prints a list of books in a formatted manner
- `membersPrinter`: Prints a list of members in a formatted manner

## Function Details

### `DisplayMenu`
- **Description**: Displays the main menu of the library management system, listing all available options in a formatted layout
- **Usage**: Called to present the user with choices for interacting with the system

### `AddBook`
- **Description**: Prompts the user for a book's title and author, creates a new `Book` struct with a unique ID, and adds it to the library using `LibraryManager.AddBook`
- **Details**: 
  - Uses `getNonEmptyInput` to ensure valid title and author inputs
  - Assigns a unique ID to the book using a global `bookCount` variable
  - Sets the initial status of the book to "Available"

### `AddMember`
- **Description**: Prompts the user for a member's name, creates a new `Member` struct with a unique ID, and adds it to the library using `LibraryManager.AddMember`
- **Details**:
  - Uses `getNonEmptyInput` to ensure a valid name is provided
  - Assigns a unique ID to the member using a global `memberCount` variable

### `RemoveBook`
- **Description**: Prompts the user for a book ID and removes the corresponding book from the library using `LibraryManager.RemoveBook`
- **Details**:
  - Uses `getValidID` to ensure a valid book ID is provided
  - If the book is borrowed, it is automatically returned before removal

### `RemoveMember`
- **Description**: Prompts the user for a member ID and removes the corresponding member from the library using `LibraryManager.RemoveMember`
- **Details**:
  - Uses `getValidID` to ensure a valid member ID is provided
  - Any books borrowed by the member are automatically returned

### `BorrowBook`
- **Description**: Prompts the user for a book ID and a member ID, then allows the member to borrow the book using `LibraryManager.BorrowBook`
- **Details**:
  - Uses `getValidID` to ensure valid book and member IDs are provided
  - Checks if the book is available before allowing the borrow operation

### `ReturnBook`
- **Description**: Prompts the user for a book ID and a member ID, then allows the member to return the book using `LibraryManager.ReturnBook`
- **Details**:
  - Uses `getValidID` to ensure valid book and member IDs are provided
  - Ensures the book is actually borrowed by the member before allowing the return

### `ListMembers`
- **Description**: Retrieves the list of all members using `LibraryManager.ListMembers` and prints them using `membersPrinter`
- **Details**:
  - If no members are registered, it informs the user

### `ListAllAvailableBooks`
- **Description**: Retrieves the list of all available books using `LibraryManager.ListAvailableBooks` and prints them using `bookPrinter`
- **Details**:
  - If no books are available, it informs the user

### `ListAllBorrowedByAMember`
- **Description**: Prompts the user for a member ID, retrieves the list of books borrowed by that member using `LibraryManager.ListBorrowedBooks`, and prints them using `bookPrinter`
- **Details**:
  - Uses `getValidID` to ensure a valid member ID is provided
  - If the member has no borrowed books, it informs the user

### `getNonEmptyInput`
- **Description**: Takes a prompt string and repeatedly asks the user for input until a non-empty string is provided
- **Details**:
  - If the user enters 'm', it returns an empty string and a boolean `false` to indicate the user wants to return to the menu
  - Used for inputs like book titles, author names, and member names

### `getValidID`
- **Description**: Takes a function (e.g., `takeBookID` or `takeMemberID`) that prompts for an ID and repeatedly calls it until a valid positive integer is provided
- **Details**:
  - If the user interrupts by entering 'm', it returns 0 and a boolean `false`
  - Used to ensure valid IDs are provided for operations like borrowing or removing books/members

### `takeBookID`
- **Description**: Prompts the user for a book ID, validates it as a positive integer, and returns the ID or an error
- **Details**:
  - Allows the user to enter 'm' to return to the menu

### `takeMemberID`
- **Description**: Prompts the user for a member ID, validates it as a positive integer, and returns the ID or an error
- **Details**:
  - Allows the user to enter 'm' to return to the menu

### `bookPrinter`
- **Description**: Takes a slice of `Book` structs and prints each book's details (ID, title, author) in a formatted manner
- **Details**:
  - Also prints the total number of books in the list

### `membersPrinter`
- **Description**: Takes a slice of `Member` structs and prints each member's details (ID, name) in a formatted manner
- **Details**:
  - Also prints the total number of members in the list

## Notes
- **Global Variables**: The system uses global variables `bookCount` and `memberCount` to generate unique IDs for books and members, respectively. In a more complex system, a better ID generation mechanism might be preferred
- **Error Handling**: Errors are handled by printing messages to the console. In a production system, more sophisticated error handling (e.g., logging) might be necessary
- **Dependencies**: This documentation assumes that the `LibraryManager` interface is implemented in `library_services.go`, which manages the underlying data structures for books and members
