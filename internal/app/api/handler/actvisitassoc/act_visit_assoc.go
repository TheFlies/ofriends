package actvisitassochandler

import (
	"context"
	"encoding/json"
	"net/http"
	"errors"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.ActVisitAssoc, error)
		GetAll(ctx context.Context) ([]types.ActVisitAssoc, error)
		GetByVisitID(ctx context.Context, visitID string) ([]types.ActVisitAssoc, error)
		GetByActID(ctx context.Context, actID string) ([]types.ActVisitAssoc, error)
		Create(ctx context.Context, actVisitAssoc types.ActVisitAssoc) (string, error)
		Update(ctx context.Context, actVisitAssoc types.ActVisitAssoc) error
		Delete(ctx context.Context, id string) error
		DeleteByVisitID(ctx context.Context, visitID string) error
	}

	// Handler is activity associate web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api activity associate handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get activity associate HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	actVisitAssoc, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, actVisitAssoc)
}

// GetActVisitAssocs handle get all activity associates HTTP Request
func (h *Handler) GetActVisitAssocs(w http.ResponseWriter, r *http.Request) {
	visitID := r.URL.Query().Get("visitid")
	var actVisitAssocs []types.ActVisitAssoc
	var err error
	if visitID != "" {
		actVisitAssocs, err = h.srv.GetByVisitID(r.Context(), visitID)
	} else {
		actVisitAssocs, err = h.srv.GetAll(r.Context())
	}
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, actVisitAssocs)
}

// GetByActID handle get activity associate by Activity ID HTTP request
func (h *Handler) GetByActID(w http.ResponseWriter, r *http.Request) {
	actVisitAssocs, err := h.srv.GetByActID(r.Context(), mux.Vars(r)["activityID"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, actVisitAssocs)
}

// Create handle insert activity associate HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var actVisitAssoc types.ActVisitAssoc

	if err := json.NewDecoder(r.Body).Decode(&actVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if id, err := h.srv.Create(r.Context(), actVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		actVisitAssoc.ID = id
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": actVisitAssoc.ID})
}

// Update handle modify activity associate HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var actVisitAssoc types.ActVisitAssoc
	if err := json.NewDecoder(r.Body).Decode(&actVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := h.srv.Update(r.Context(), actVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": actVisitAssoc.ID, "status": "success"})
}

// Delete handle delete activity associate HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// DeleteByVisitID handle delete activity associate HTTP Request
func (h *Handler) DeleteByVisitID(w http.ResponseWriter, r *http.Request) {
	visitID := r.URL.Query().Get("visitid")
	if visitID == "" {
		respond.Error(w, errors.New("invalid Request"), http.StatusInternalServerError)
		return
	}
	if err := h.srv.DeleteByVisitID(r.Context(), visitID); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
