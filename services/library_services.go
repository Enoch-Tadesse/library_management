package services

import (
	"fmt"
	"library_management/models"
	"slices"
)

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(member models.Member)
	RemoveBook(bookID int)
	RemoveMember(memberID int) error
	ListMembers() []models.Member
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
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

	for _, book := range member.BorrowedBooks {
		if b, ok := l.Books[book.ID]; ok {
			b.Status = "Available"
			l.Books[book.ID] = b // Save the updated book back
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
	if book.Status != "Available" {
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

	if book.Status != "Available" {
		return fmt.Errorf("book %s is not available", book.Title)
	}

	// set the status to Borrowed
	book.Status = "Borrowed"
	fmt.Println("here")
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

	if book.Status == "Available" {
		return fmt.Errorf("book %s is not borrowed", book.Title)
	}

	idx := 0
	for idx < len(member.BorrowedBooks) {
		borrowdBook := member.BorrowedBooks[idx]
		if borrowdBook.ID == book.ID {
			break
		}
		idx += 1
	}

	if idx == len(member.BorrowedBooks) {
		return fmt.Errorf("book %s is not borrowed by %s", book.Title, member.Name)
	}

	book.Status = "Available"
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
	for _, book := range l.Books {
		if book.Status == "Available" {
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
