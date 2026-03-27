package tasks

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/neptune2k21/chemin-d-or/internal/domain"
	"github.com/neptune2k21/chemin-d-or/pkg/response"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "corps de la requête invalide")
		return
	}

	task, err := h.svc.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, task)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, err := h.svc.GetByID(r.Context(), id)
	if errors.Is(err, ErrTaskNotFound) {
		response.Error(w, http.StatusNotFound, "tâche non trouvée")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "erreur interne du serveur")
		return
	}
	response.JSON(w, http.StatusOK, task)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.svc.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "erreur interne du serveur")
		return
	}
	if tasks == nil {
		tasks = []*domain.Task{}
	}
	response.JSON(w, http.StatusOK, tasks)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "corps de la requête invalide")
		return
	}

	task, err := h.svc.Update(r.Context(), id, req)
	if errors.Is(err, ErrTaskNotFound) {
		response.Error(w, http.StatusNotFound, "tâche non trouvée")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "erreur interne du serveur")
		return
	}
	response.JSON(w, http.StatusOK, task)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.svc.Delete(r.Context(), id)
	if errors.Is(err, ErrTaskNotFound) {
		response.Error(w, http.StatusNotFound, "task not found")
		return
	}
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "erreur interne du serveur")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
