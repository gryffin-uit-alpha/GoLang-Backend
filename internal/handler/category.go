package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gryffin-uit-alpha/GoLang-Backend/internal/db"
	"github.com/gryffin-uit-alpha/GoLang-Backend/internal/utils"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func CreateCategory(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	slug := utils.GenerateSlug(req.Name)

	category, err := queries.CreateCategory(r.Context(), db.CreateCategoryParams{
		Name: req.Name,
		Slug: slug,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func ListCategories(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	categories, err := queries.ListCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
