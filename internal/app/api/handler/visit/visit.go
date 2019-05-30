package visithandler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Visit, error)
		GetByCustomerID(ctx context.Context, customerId string) ([]types.Visit, error)
		GetAll(ctx context.Context) ([]types.Visit, error)
		Create(ctx context.Context, visit types.Visit) (string, error)
		Update(ctx context.Context, visit types.Visit) error
		Delete(ctx context.Context, id string) error
	}
	MailService interface {
		Send(ctx context.Context, now time.Time) error
		SendEmailByVisitID(ctx context.Context, visitID string) error
	}

	// Handler is visit web handler
	Handler struct {
		srv     service
		logger  glog.Logger
		mailSrv MailService
	}
)

// New return new rest api visit handler
func New(s service, mSrv MailService, l glog.Logger) *Handler {
	return &Handler{
		srv:     s,
		logger:  l,
		mailSrv: mSrv,
	}
}

// Get handle get visit HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	visit, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, visit)
}

// GetByCustomerID handle get visits HTTP Request by customerID
func (h *Handler) GetByCustomerID(w http.ResponseWriter, r *http.Request) {
	visits, err := h.srv.GetByCustomerID(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, visits)

}

// GetAll handle get visits HTTP Request
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	visits, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, http.StatusOK, visits)

}

// Create handle insert visit HTTP Request
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var visit types.Visit

	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if id, err := h.srv.Create(r.Context(), visit); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	} else {
		visit.ID = id
	}

	respond.JSON(w, http.StatusCreated, map[string]string{"id": visit.ID})
}

// Update handle modify visit HTTP Request
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var visit types.Visit

	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		if err == io.EOF {
			respond.Error(w, errors.New("Invalid request method"), http.StatusMethodNotAllowed)
			return
		}
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if err := h.srv.Update(r.Context(), visit); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": visit.ID})
}

// Delete handle delete visit HTTP Request
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.srv.Delete(r.Context(), mux.Vars(r)["id"]); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func (h *Handler) SendMail(w http.ResponseWriter, r *http.Request) {
	visitID := mux.Vars(r)["visitID"]
	err := h.mailSrv.SendEmailByVisitID(r.Context(), visitID)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success", "message": "the email have been seen to IC department"})
	return

}
