package services

import (
	"fmt"
	"library_management/models"
	"slices"
)

// LibraryManager defines the interface for managing books and members in a library.
type LibraryManager interface {
	// AddBook adds a book to the library's collection.
	AddBook(book models.Book)

	// AddMember registers a new member.
	AddMember(member models.Member)

	// RemoveBook deletes a book by ID, returning it first if borrowed.
	RemoveBook(bookID int)

	// RemoveMember removes a member by ID and returns all their borrowed books.
	RemoveMember(memberID int) error

	// ListMembers returns all registered library members.
	ListMembers() []models.Member

	// BorrowBook allows a member to borrow a book.
	BorrowBook(bookID int, memberID int) error

	// ReturnBook allows a member to return a borrowed book.
	ReturnBook(bookID int, memberID int) error

	// ListAvailableBooks returns all books currently available for borrowing.
	ListAvailableBooks() []models.Book

	// ListBorrowedBooks returns all books borrowed by a member.
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]*models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*models.Book),
		Members: make(map[int]*models.Member),
	}
}

// ListMembers fetches  all the members
// from Members map
func (l *Library) ListMembers() []models.Member {
	var allMember []models.Member
	for _, v := range l.Members {
		allMember = append(allMember, *v)
	}
	return allMember
}

// RemoveMember removes a member from
// Members map
func (l *Library) RemoveMember(memberID int) error {
	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with id %d does not exist", memberID)
	}

	// Before removing a member, return all
	// the borrowed books by the member ID
	for _, book := range member.BorrowedBooks {
		if b, ok := l.Books[book.ID]; ok {
			b.Status = models.Available
		}
	}

	delete(l.Members, memberID)
	return nil
}

// AddBook inserts a book
// into Books map
func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = &book
}

// RemoveBook removes a book
// from Books map
func (l *Library) RemoveBook(bookID int) {
	book, ok := l.Books[bookID]
	if !ok {
		return
	}
	// Before removing the book, if the book was borrowed
	// make the member return the book
	if book.Status != models.Available {
		for _, member := range l.Members {
			err := l.ReturnBook(bookID, member.ID)
			if err == nil {
				break
			}
		}
	}
	delete(l.Books, bookID)
}

// BorrowBook inserts a book into member's
// borrowed collection
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id %d does not exist", bookID)
	}
	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberID)
	}

	// Book is not available
	if book.Status != models.Available {
		return fmt.Errorf("book %s is not available", book.Title)
	}

	// set the status to Borrowed
	book.Status = models.Borrowed
	member.BorrowedBooks = append(member.BorrowedBooks, *book)
	return nil
}

// ReturnBook removes a book from member's borrowed collection
// if it does not exist, it returns an error
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.Books[bookID]
	if !ok {
		return fmt.Errorf("book with id %d does not exist", bookID)
	}
	member, ok := l.Members[memberID]
	if !ok {
		return fmt.Errorf("member with id %d not found", memberID)
	}

	if book.Status == models.Available {
		return fmt.Errorf("book %s is not borrowed", book.Title)
	}

	// Find the book in the borrowed books slice
	idx := 0
	for idx < len(member.BorrowedBooks) {
		borrowdBook := member.BorrowedBooks[idx]
		if borrowdBook.ID == book.ID {
			break
		}
		idx += 1
	}

	// Book was never found
	if idx == len(member.BorrowedBooks) {
		return fmt.Errorf("book %s is not borrowed by %s", book.Title, member.Name)
	}

	// Set the book to Available
	book.Status = models.Available
	member.BorrowedBooks = slices.Delete(member.BorrowedBooks, idx, idx+1)

	return nil
}

// AddMember adda a member into
// members map
func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = &member
}

// ListAvailableBooks fetches and return a slice
// of all available books from books map
func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	// Iterate over all the books and
	// accumulate available books
	for _, book := range l.Books {
		if book.Status == models.Available {
			availableBooks = append(availableBooks, *book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks returns all the borrowed books
// by a specific member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, ok := l.Members[memberID]
	if !ok {
		fmt.Printf("member with id %d does not exist\n", memberID)
		return nil
	}

	return member.BorrowedBooks
}
