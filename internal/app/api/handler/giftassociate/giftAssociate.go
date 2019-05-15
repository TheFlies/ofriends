package giftassociatehandler

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
		Get(ctx context.Context, id string) (*types.GiftAssociate, error)
		GetAll(ctx context.Context) ([]types.GiftAssociate, error)
		GetByVisitIDNCustomerID(ctx context.Context, visitID string, customerID string) ([]types.GiftAssociate, error)
		Create(ctx context.Context, giftAssociate types.GiftAssociate) (string, error)
		Update(ctx context.Context, giftAssociate types.GiftAssociate) error
		Delete(ctx context.Context, id string) error
		DeleteByVisitIDNCustomerID(ctx context.Context, visitID string, customerID string) error
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
	giftAssociate, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, giftAssociate)
}

// GetAll handle get all gift associates HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	giftAssociates, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, giftAssociates)
}

// GetByVisitIDNCustomerID handle get gift associate by Visit ID HTTP request
func (h *Handler) GetByVisitIDNCustomerID(w http.ResponseWriter, r *http.Request) {
	giftAssociates, err := h.srv.GetByVisitIDNCustomerID(r.Context(), mux.Vars(r)["visitID"], mux.Vars(r)["customerID"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, giftAssociates)
}

// Create handle insert gift associate HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var giftAssociate types.GiftAssociate

	if err := json.NewDecoder(r.Body).Decode(&giftAssociate); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if id, err := h.srv.Create(r.Context(), giftAssociate); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		giftAssociate.ID = id
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": giftAssociate.ID})
}

// Update handle modify gift associate HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var giftAssociate types.GiftAssociate
	if err := json.NewDecoder(r.Body).Decode(&giftAssociate); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := h.srv.Update(r.Context(), giftAssociate); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": giftAssociate.ID, "status": "success"})
}

// Delete handle delete gift associate HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// DeleteByVisitIDNCustomerID handle delete gift associate HTTP Request
func (h *Handler) DeleteByVisitIDNCustomerID(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.DeleteByVisitIDNCustomerID(r.Context(), mux.Vars(r)["visitID"], mux.Vars(r)["customerID"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
