package web

import (
	"booksAdmin/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}

// GET /books.
func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.service.GetBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get Books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)

}

// GET /books/search
func (h *BookHandlers) GetBookByTitle(writer http.ResponseWriter, request *http.Request) {

	title := request.URL.Query().Get("title")
	books, err := h.service.SearchBookByName(title)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Failed to get Books", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(books)
}

// POST /books.
func (h *BookHandlers) CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book service.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		fmt.Println("Error decode", err)
		http.Error(writer, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	fmt.Println(book)
	createdBook, err := h.service.CreateBook(&book)
	if err != nil {
		fmt.Println("Error create book", err)
		http.Error(writer, "Failed to create book", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(&createdBook)
}

// GET /books/{id}.
func (h *BookHandlers) GetBookByID(writer http.ResponseWriter, request *http.Request) {
	// idStr := request.URL.Query().Get("id")
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book, err := h.service.GetBook(id)
	if err != nil {
		http.Error(writer, "Failed to get book", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(writer, "Book not found", http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&book)
}

// PUT /books/{id}.
func (h *BookHandlers) UpdateBook(writer http.ResponseWriter, request *http.Request) {
	// idStr := request.URL.Query().Get("id")
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	var book service.Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		http.Error(writer, "Invalid Request Payload", http.StatusBadRequest)
		return
	}
	book.ID = id
	if err := h.service.UpdateBook(&book); err != nil {
		http.Error(writer, "Failed to update book", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&book)
}

// DELETE /books/{id}.
func (h *BookHandlers) DeleteBook(writer http.ResponseWriter, request *http.Request) {
	// idStr := request.URL.Query().Get("id")
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		http.Error(writer, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// POST /books/simulate-reading
func (h *BookHandlers) SimulateReading(writer http.ResponseWriter, request *http.Request) {
	durationStr := request.URL.Query().Get("duration")
	duration, err := strconv.Atoi(durationStr) // Converte para inteiro
	if err != nil {
		http.Error(writer, "Invalid duration", http.StatusBadRequest)
		return
	}
	type payload struct {
		BookIDs []int `json:"book_ids"`
	}

	var booksIDs payload
	err = json.NewDecoder(request.Body).Decode(&booksIDs)
	if err != nil {
		fmt.Println("Error decode", err)
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}
	books := h.service.SimulateReadingParallel(booksIDs.BookIDs, time.Duration(duration)*time.Second)

	if len(books) == 0 {
		http.Error(writer, "No books found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&books)
}
