package friendhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"errors"
	"io"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Friend, error)
		GetAll(ctx context.Context) ([]types.Friend, error)
		Create(ctx context.Context, gift types.Friend) error
		Update(ctx context.Context, gift types.Friend) error
		Delete(ctx context.Context, id string) error
	}

	// Handler is friend web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api friend handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get friend HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	friend, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, friend)
}

// GetAll handle get gifts HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	friends, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, friends)

}

// Create handle insert gift HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var friend types.Friend

	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if err := h.srv.Create(r.Context(), friend); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusCreated, map[string]string{"id": friend.ID})
}

// Update handle modify gift HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var friend types.Friend

	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if err := h.srv.Update(r.Context(), friend); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": friend.ID})
}

// Delete handle delete gift HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}
