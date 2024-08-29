package cli

import (
	"booksAdmin/internal/service"
	"fmt"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *service.BookService
}

// Função construtora pra gerar o Struct
func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go books <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <title>")
			return
		}
		bookName := os.Args[2]
		cli.SearchBook(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
			return
		}
		bookIDs := os.Args[2:] // Pega todos os argumentos depois do primeiro
		cli.simulateReadingOS(bookIDs)

	default:
		fmt.Println("Invalid command")
	}

}

func (cli *BookCLI) SearchBook(title string) {
	books, err := cli.service.SearchBookByName(title)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%d Books found\n", len(books))
	for _, book := range books {
		fmt.Printf("ID: %d - Book: %s - %s - %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReadingOS(bookIDsStr []string) {
	var bookIDs []int
	for _, idStr := range bookIDsStr {
		id, err := strconv.Atoi(idStr) // Converte para inteiro
		if err != nil {
			fmt.Println("Invalid book ID:", idStr)
			continue
		}
		bookIDs = append(bookIDs, id)
	}

	results := cli.service.SimulateReadingParallel(bookIDs, 2*time.Second)
	for _, result := range results {
		fmt.Println(result)
	}

}
