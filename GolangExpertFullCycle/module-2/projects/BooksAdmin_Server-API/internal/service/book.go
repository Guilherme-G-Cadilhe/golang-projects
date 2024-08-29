package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    genre varchar(255) NOT NULL
);
	`
	_, err := db.Exec(query)
	return err
}

func (s *BookService) CreateBook(book *Book) (Book, error) {
	query := "INSERT INTO books (title, author, genre) VALUES (?, ?, ?)"

	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {

		return Book{}, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return Book{}, err
	}
	book.ID = int(lastInsertID)
	return *book, nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT id, title, author, genre FROM books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBook(id int) (*Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := s.db.Exec(query, id)
	return err
}

func (s *BookService) SearchBookByName(name string) ([]Book, error) {
	query := "SELECT id, title, author, genre FROM books WHERE title LIKE ?"
	rows, err := s.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) SimulateReading(id int, duration time.Duration, results chan<- string) {
	book, err := s.GetBook(id)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Book %d not found", id)
		return
	}
	time.Sleep(duration)
	results <- fmt.Sprintf("Book %s read successfully", book.Title)
}

func (s *BookService) SimulateReadingParallel(bookIDs []int, duration time.Duration) []string {
	results := make(chan string, len(bookIDs)) // Canal com buffer

	for _, id := range bookIDs {
		go func(bookId int) {
			s.SimulateReading(bookId, duration, results)
		}(id) // Se auto inicializa
	}

	var responses []string
	// Fica lendo indefinidamente
	// for res := range results {
	// 	responses = append(responses, res)
	// }

	// Aguarda todos terminarem
	// Inicializa um, aguarda o primeiro resultado do canal, inicializa outro, aguarda o outro, e etc, ate o canal estiver vazio ( Ter percorrido todo o Length do buffer )
	for range bookIDs {
		responses = append(responses, <-results) // Pausa enquanto aguarda
	}

	close(results)
	return responses

}
