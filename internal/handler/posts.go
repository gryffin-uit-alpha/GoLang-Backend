package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/gryffin-uit-alpha/GoLang-Backend/internal/db"
	"github.com/gryffin-uit-alpha/GoLang-Backend/internal/utils"
)

type CreatePostRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   string `json:"author_id"`
	CategoryID string `json:"category_id"`
	Status     string `json:"status"`
}

func CreatePost(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	slug := utils.GenerateSlug(req.Title)
	authorUUID, _ := uuid.Parse(req.AuthorID)

	var categoryUUID uuid.NullUUID
	if req.CategoryID != "" {
		parsed, _ := uuid.Parse(req.CategoryID)
		categoryUUID = uuid.NullUUID{UUID: parsed, Valid: true}
	}

	post, err := queries.CreatePost(r.Context(), db.CreatePostParams{
		Title:      req.Title,
		Slug:       slug,
		Content:    req.Content,
		AuthorID:   authorUUID,
		CategoryID: categoryUUID,
		Status:     req.Status,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func ListPosts(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	posts, err := queries.ListPublishedPosts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPost(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	post, err := queries.GetPostBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func DeletePost(queries *db.Queries, w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	postID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = queries.SoftDeletePost(r.Context(), postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
