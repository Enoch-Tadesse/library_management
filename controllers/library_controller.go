package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"strings"
)

type menuItem struct {
	Label  string
	Action func(l services.LibraryManager)
}

var Menu = map[string]menuItem{
	"1": {"Add a new book", AddBook},
	"2": {"Add a new member", AddMember},
	"3": {"Remove an existing member", RemoveMember},
	"4": {"List all members", ListMembers},
	"5": {"Remove an existing book", RemoveBook},
	"6": {"Borrow a book", BorrowBook},
	"7": {"Return a book", ReturnBook},
	"8": {"List all available books", ListAllAvailableBooks},
	"9": {"List all borrowed books by a member", ListAllBorrowedByAMember},
}

// DisplayMenu displays the menu in appealing
// user interface
func DisplayMenu() {
	width := 40
	fmt.Println(strings.Repeat("=", width))

	// try to center the title
	title := "Library Management Menu"
	padding := max(((width - len(title)) / 2), 0)

	fmt.Println(strings.Repeat(" ", padding), title)
	fmt.Println(strings.Repeat("=", width))

	for i := 1; i <= len(Menu); i++ {
		key := fmt.Sprintf("%d", i)
		fmt.Printf(" %d. %-35s\n", i, Menu[key].Label)
	}

	fmt.Println(strings.Repeat("-", width))
	fmt.Println(" q. Exit")
	fmt.Println(strings.Repeat("=", width))
}

// keep count of how many meber there is
var memberCount = 0

// ListMembers return a list of active members 
func ListMembers(l services.LibraryManager) {
	members := l.ListMembers()

	if len(members) == 0 {
		fmt.Println("No members yet")
		return
	}
	membersPrinter(members)
}

// RemoveMember removes member by ID
func RemoveMember(l services.LibraryManager) {
	memberID, ok := getValidID(takeMemberID)
	if !ok {
		return
	}

	err := l.RemoveMember(memberID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("member removed successfully")
}

// AddMember adds a member by ID
func AddMember(l services.LibraryManager) {
	name, ok := getNonEmptyInput("Member Name (or 'm' to return): ")

	if !ok {
		return
	}
	var member models.Member
	member.Name = name
	memberCount += 1
	member.ID = memberCount
	l.AddMember(member)
	fmt.Println("member added successfully")
}

// ListAllBorrowedByAMember lists all borrowed 
// books by a specific member
func ListAllBorrowedByAMember(l services.LibraryManager) {
	memberID, ok := getValidID(takeMemberID)
	if !ok {
		return
	}
	books := l.ListBorrowedBooks(memberID)
	bookPrinter(books)
}

// ListAllAvailableBooks lists all books 
// that are not borrowed
func ListAllAvailableBooks(l services.LibraryManager) {
	books := l.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("no available books")
		return
	}

	bookPrinter(books)
}

// ReturnBook returns a book by
// member ID and book ID
func ReturnBook(l services.LibraryManager) {
	bookID, ok := getValidID(takeBookID)
	if !ok {
		return
	}
	memberID, ok := getValidID(takeMemberID)
	if !ok {
		return
	}
	err := l.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Book Returned Successfully")
}

// BorrowBook lets a member boorow a book
// by book ID
func BorrowBook(l services.LibraryManager) {

	bookID, ok := getValidID(takeBookID)
	if !ok {
		return
	}
	memberID, ok := getValidID(takeMemberID)
	if !ok {
		return
	}
	err := l.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Book Borrowed Successfully")
}

// RemoveBook removes a book from collection
// by book ID
func RemoveBook(l services.LibraryManager) {

	bookID, ok := getValidID(takeBookID)
	if !ok {
		return
	}
	l.RemoveBook(bookID)
	fmt.Println("Book Removed Successfully")
}

// keeps track of how many book entered in
// the collection
var bookCount = 0

// AddBook adds a book to the collection
func AddBook(l services.LibraryManager) {
	var book models.Book
	title, ok := getNonEmptyInput("Book Title (or 'm' to return): ")
	if !ok {
		return
	}
	book.Title = title

	author, ok := getNonEmptyInput("Author Name (or 'm' to return): ")
	if !ok {
		return
	}
	book.Author = author

	// set the status Available initially
	book.Status = "Available"
	bookCount += 1
	book.ID = bookCount

	l.AddBook(book)
	fmt.Println("Book Added Successfully")
}
