package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bhushan9001/GO_CRUD/config"
	"github.com/Bhushan9001/GO_CRUD/internal/models"
	"github.com/Bhushan9001/GO_CRUD/utils"
	"github.com/gorilla/mux"
	
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result := config.DB.Create(&book)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, book)
}

// GetBooks handles GET requests to retrieve all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Books
	result := config.DB.Find(&books)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	utils.RespondJSON(w, http.StatusOK, books)
}

// GetBook handles GET requests to retrieve a specific book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Books
	result := config.DB.First(&book, id)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, book)
}

// UpdateBook handles PUT requests to update a specific book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Books
	result := config.DB.First(&book, id)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusNotFound, "Book not found")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	config.DB.Save(&book)
	utils.RespondJSON(w, http.StatusOK, book)
}

// DeleteBook handles DELETE requests to remove a specific book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	result := config.DB.Delete(&models.Books{}, id)
	if result.Error != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	if result.RowsAffected == 0 {
		utils.RespondJSON(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, "Book successfully deleted")
}
