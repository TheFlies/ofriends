package activityvisitassociatehandler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.ActivityVisitAssociate, error)
		GetAll(ctx context.Context) ([]types.ActivityVisitAssociate, error)
		GetByVisitID(ctx context.Context, visitID string) ([]types.ActivityVisitAssociate, error)
		Create(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) (string, error)
		Update(ctx context.Context, activityVisitAssociate types.ActivityVisitAssociate) error
		Delete(ctx context.Context, id string) error
		DeleteByVisitID(ctx context.Context, visitID string) error
	}

	// Handler is gift associate web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api gift associate handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get gift associate HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	activityVisitAssociate, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, activityVisitAssociate)
}

// GetAll handle get all gift associates HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	activityVisitAssociates, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, activityVisitAssociates)
}

// GetByVisitID handle get gift associate by Visit ID HTTP request
func (h *Handler) GetByVisitID(w http.ResponseWriter, r *http.Request) {
	activityVisitAssociates, err := h.srv.GetByVisitID(r.Context(), mux.Vars(r)["visitID"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, activityVisitAssociates)
}

// Create handle insert gift associate HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var activityVisitAssociate types.ActivityVisitAssociate

	if err := json.NewDecoder(r.Body).Decode(&activityVisitAssociate); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if id, err := h.srv.Create(r.Context(), activityVisitAssociate); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		activityVisitAssociate.ID = id
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": activityVisitAssociate.ID})
}

// Update handle modify gift associate HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var activityVisitAssociate types.ActivityVisitAssociate
	if err := json.NewDecoder(r.Body).Decode(&activityVisitAssociate); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := h.srv.Update(r.Context(), activityVisitAssociate); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": activityVisitAssociate.ID, "status": "success"})
}

// Delete handle delete gift associate HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// DeleteByVisitID handle delete gift associate HTTP Request
func (h *Handler) DeleteByVisitID(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.DeleteByVisitID(r.Context(), mux.Vars(r)["visitID"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
