 package timeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/customer"
	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/giftassociate"
	"github.com/TheFlies/ofriends/internal/app/gift"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/app/visit"
	"github.com/TheFlies/ofriends/internal/app/activity"
	"github.com/TheFlies/ofriends/internal/app/actvisitassoc"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
)

type (
	Service struct {
		VisitRepo    			visit.Repository
		CustomerRepo 			customer.Repository
		CustomerVisitAssRepo    cusvisitassoc.Repository
		GiftAssRepo 			giftassociate.Repository
		GiftRepo             	gift.Repository
		ActVisitAssRepo			actvisitassoc.Repository
		ActivityRepo 			activity.Repository
		logger       			glog.Logger
	}
)

func NewService(visitRepo visit.Repository, customerRepo customer.Repository, customeVisitAssRepo cusvisitassoc.Repository,
				giftAssRepo giftassociate.Repository, giftRepo gift.Repository,
				actVisitAssRepo actvisitassoc.Repository, activityRepo activity.Repository,
				l glog.Logger) *Service {
	return &Service{
		VisitRepo:    			visitRepo,
		CustomerRepo: 			customerRepo,
		CustomerVisitAssRepo:   customeVisitAssRepo,
		GiftAssRepo: 			giftAssRepo,
		GiftRepo:               giftRepo,
		ActVisitAssRepo:		actVisitAssRepo,
		ActivityRepo:           activityRepo,
		logger:       			l,
	}
}
func (m *Service) FindTimelineByDay(ctx context.Context, dayTime int64) ([]types.Timeline, error) {
	m.logger.Debugf("Get visit before day %d", dayTime)
	var timelines []types.Timeline
	visits, err := m.VisitRepo.FindInCommingVisit(ctx, dayTime)
	if err != nil {
		m.logger.Errorf("%v", err)
		return nil, errors.Wrap(err, "can't get in comming visit")
	}
	if (len(visits) == 0){
		m.logger.Debugf("Dont have any visits")
	}
	for _, visit := range visits {
		var timeline types.Timeline
		timeline.Visit = visit

		// Get all customer in this visit
		assList, err := m.CustomerVisitAssRepo.FindByVisitID(ctx, visit.ID)
		if err != nil {
			m.logger.Errorf("%v", err)
			return nil, errors.Wrap(err, "can't get Customer Visit Associating")
		}
		for _, ass := range assList {
			customer, err := m.CustomerRepo.FindByID(ctx, ass.CustomerID)
			if err != nil {
				m.logger.Errorf("%v", err)
				return nil, errors.Wrap(err, "can't get Customer")
			}
			timeline.Customers = append(timeline.Customers, customer)

			// Get all gift in this visit
			giftAssList, err := m.GiftAssRepo.FindByCusVisitAssocID(ctx, ass.ID)
			for _, giftAss := range giftAssList {
				gift, err := m.GiftRepo.FindByID(ctx, giftAss.GiftID)
				if err != nil {
					m.logger.Errorf("%v", err)
					return nil, errors.Wrap(err, "can't get gift")
				}
				timeline.Gifts = append(timeline.Gifts, gift)
			}
		}

		// Get all activity in this visit
		actAssList, err := m.ActVisitAssRepo.FindByVisitID(ctx, visit.ID)
		if err != nil {
			m.logger.Errorf("%v", err)
			return nil, errors.Wrap(err, "can't get Customer Visit Associating")
		}
		for _, ass := range actAssList {
			activity, err := m.ActivityRepo.FindByID(ctx, ass.ActivityID)
			if err != nil {
				m.logger.Errorf("%v", err)
				return nil, errors.Wrap(err, "can't get Customer")
			}
			timeline.Activitys = append(timeline.Activitys, activity)
		}
		timelines = append(timelines, timeline)

	}
	
	return timelines, nil
}
