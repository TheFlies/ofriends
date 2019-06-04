package timelinehandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
	"github.com/gorilla/mux"
)

type (
	service interface {
		FindTimelineByDay(ctx context.Context, dayTime int64) ([]types.Timeline, error)
	}

	// Handler is timeline web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api timeline handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}


// GetAll handle get customer HTTP Request
func (h *Handler) FindTimelineByDay(w http.ResponseWriter, r *http.Request) {
	daytime, _ := strconv.ParseInt(mux.Vars(r)["daytime"], 10, 64)
	h.logger.Infof("daytime : %v", daytime)
	timelines, err := h.srv.FindTimelineByDay(r.Context(), daytime)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, timelines)

}
