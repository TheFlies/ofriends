package customerhandler

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
		Get(ctx context.Context, id string) (*types.Customer, error)
		GetAll(ctx context.Context) ([]types.Customer, error)
		Create(ctx context.Context, customer types.Customer) (string, error)
		Update(ctx context.Context, customer types.Customer) error
		Delete(ctx context.Context, id string) error
	}

	// Handler is customer web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api customer handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get customer HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	customer, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, customer)
}

// GetAll handle get customer HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, customers)

}

// Create handle insert customer HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var customer types.Customer
	
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close() 

	if id, err := h.srv.Create(r.Context(), customer); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		customer.ID = id
	}

	respond.JSON(w, http.StatusCreated, map[string]string{"id": customer.ID})
}

// Update handle modify customer HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var customer types.Customer

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	defer r.Body.Close() 

	if err := h.srv.Update(r.Context(), customer); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": customer.ID})
}

// Delete handle delete customer HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
