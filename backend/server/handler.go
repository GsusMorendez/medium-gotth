package server

import (
	"net/http"
	"version1-medium-gotth/backend/repository"
)

type Handler struct {
	Painter     *Painter
	DataManager *repository.DataManager
}

func NewHandler(painter *Painter, dataManager *repository.DataManager) *Handler {
	return &Handler{painter, dataManager}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	h.Painter.HelloWorld(w, r.Context())
}

func (h *Handler) MainView(w http.ResponseWriter, r *http.Request) {
	data := h.DataManager.GetSampleData()
	h.Painter.MainPage(w, r.Context(), data)
}
