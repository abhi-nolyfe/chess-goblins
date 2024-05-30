package handlers

import (
	"net/http"

	"github.com/abhi-nolyfe/chess-goblins/internal/templates"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}