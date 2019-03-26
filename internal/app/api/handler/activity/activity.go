package activityhandler

import (
	"context"
	"net/http"
	"encoding/json"

	"errors"
	"io"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Activity, error)
		GetByVisitID(ctx context.Context, visitId string) ([]types.Activity, error)
		GetAll(ctx context.Context) ([]types.Activity, error)
		Create(ctx context.Context, visit types.Activity) (string, error)
		Update(ctx context.Context, visit types.Activity) error
		Delete(ctx context.Context, id string) error
	}

	// Handler is visit web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api visit handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get activity HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	act, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, act)
}

// GetByFriendID handle get activities HTTP Request by visitID
func (h *Handler) GetByVisitID(w http.ResponseWriter, r *http.Request) {
	acts, err := h.srv.GetByVisitID(r.Context(),mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, acts)

}

// GetAll handle get activities HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	acts, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, acts)

}

// Create handle insert activity HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var act types.Activity
	
	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close() 

	if id, err := h.srv.Create(r.Context(), act); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		act.ID = id
	}

	respond.JSON(w, http.StatusCreated, map[string]string{"id": act.ID})
}

// Update handle modify actitity HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var act types.Activity

	if err := json.NewDecoder(r.Body).Decode(&act); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	defer r.Body.Close() 

	if err := h.srv.Update(r.Context(), act); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": act.ID})
}

// Delete handle delete activity HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
