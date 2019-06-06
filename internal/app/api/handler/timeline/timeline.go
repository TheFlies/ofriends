package timelinehandler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/respond"
)

type (
	service interface {
		FindTimelineByDay(ctx context.Context, dayTime int64) ([]types.Timeline, error)
		FindTimelineInRange(ctx context.Context, startTime, endTime int64) ([]types.Timeline, error)
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


// FindTimelineByDay handle get timeline by day Or get visit in input range of time in HTTP Request
func (h *Handler) FindTimelineByDay(w http.ResponseWriter, r *http.Request) {
	daytime,_ := strconv.ParseInt(r.URL.Query().Get("daytime"),10,64)
	starttime,_:= strconv.ParseInt(r.URL.Query().Get("starttime"),10,64)
	endtime,_:= strconv.ParseInt(r.URL.Query().Get("endtime"),10,64)

	var timelines []types.Timeline
	var err error

	if daytime != 0 {
		timelines, err = h.srv.FindTimelineByDay(r.Context(), daytime)
	} else {
		timelines, err = h.srv.FindTimelineInRange(r.Context(), starttime, endtime)
	}
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	
	respond.JSON(w, http.StatusOK, timelines)
}
