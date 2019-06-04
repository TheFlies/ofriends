package cusvisitassochandler

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
		Get(ctx context.Context, id string) (*types.CusVisitAssoc, error)
		GetAll(ctx context.Context) ([]types.CusVisitAssoc, error)
		GetByVisitID(ctx context.Context, visitID string) ([]types.CusVisitAssoc, error)
		Create(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) (string, error)
		Update(ctx context.Context, cusVisitAssoc types.CusVisitAssoc) error
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
	cusVisitAssoc, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, cusVisitAssoc)
}

// GetAll handle get all gift associates HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	var cusVisitAssocs []types.CusVisitAssoc
	var err error
	if visitID := queryValues.Get("visitid"); visitID != "" {
		cusVisitAssocs, err = h.srv.GetByVisitID(r.Context(), visitID)
	} else {
		cusVisitAssocs, err = h.srv.GetAll(r.Context())
	}

	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, cusVisitAssocs)
}

// GetByVisitID handle get gift associate by Visit ID HTTP request
func (h *Handler) GetByVisitID(w http.ResponseWriter, r *http.Request) {
	cusVisitAssocs, err := h.srv.GetByVisitID(r.Context(), mux.Vars(r)["visitID"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, cusVisitAssocs)
}

// Create handle insert gift associate HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var cusVisitAssoc types.CusVisitAssoc

	if err := json.NewDecoder(r.Body).Decode(&cusVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if id, err := h.srv.Create(r.Context(), cusVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		cusVisitAssoc.ID = id
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": cusVisitAssoc.ID})
}

// Update handle modify gift associate HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var cusVisitAssoc types.CusVisitAssoc
	if err := json.NewDecoder(r.Body).Decode(&cusVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := h.srv.Update(r.Context(), cusVisitAssoc); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": cusVisitAssoc.ID, "status": "success"})
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
	queryValues := r.URL.Query()
	if err := h.srv.DeleteByVisitID(r.Context(), queryValues.Get("visitid")); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
