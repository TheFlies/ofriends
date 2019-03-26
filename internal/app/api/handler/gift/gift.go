package gifthandler

import (
	"context"
	"encoding/json"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"net/http"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Gift, error)
		GetAll(ctx context.Context) ([]types.Gift, error)
		Create(ctx context.Context, gift types.Gift) (string, error)
		Update(ctx context.Context, gift types.Gift) error
		Delete(ctx context.Context, id string) error
	}

	// Handler is gift web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api gift handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get gift HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	gift, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, gift)
}

// GetAll handle get gifts HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	gifts, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, gifts)
}

// Create handle insert gift HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var gift types.Gift

	if err := json.NewDecoder(r.Body).Decode(&gift); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if id, err := h.srv.Create(r.Context(), gift); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		gift.ID = id
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": gift.ID})
}

// Update handle modify gift HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var gift types.Gift
	if err := json.NewDecoder(r.Body).Decode(&gift); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := h.srv.Update(r.Context(), gift); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": gift.ID, "status":"success"})
}

// Delete handle delete gift HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
