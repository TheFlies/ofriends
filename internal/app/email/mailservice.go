package email

import (
	"github.com/TheFlies/ofriends/internal/app/customer"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/app/visit"
	envconfig "github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/pkg/errors"
	"time"
)

type (
	MailService interface {
		Send(data types.NotificationData) error
	}

	SendMailService struct {
		VisitRepo    visit.Repository
		CustomerRepo customer.Repository
		logger       glog.Logger
	}

	TimeConf struct {
		/*
			the service will send email to notify internal communication the visit is coming  (day)
			default is 7days = one week before visit come
		*/
		SendingCondition int    `envconfig:"SENDING_CONDITION" default:"7"`
		TimeZoneLocation string `envconfig:"TIME_ZONE_LOCATION" default:"Asia/Bangkok"`
	}
)

func NewSendMailService(visitRepo visit.Repository, customerRepo customer.Repository, l glog.Logger) *SendMailService {
	return &SendMailService{
		VisitRepo:    visitRepo,
		CustomerRepo: customerRepo,
		logger:       l,
	}
}
func (m *SendMailService) Send(data types.NotificationData) error {
	var conf TimeConf
	envconfig.Load(&conf)
	yearNow, monthNow, dayNow := time.Time.Date(time.Now())
	//dm,mm,ym := time.Time.Date(time.Now())
	monthNow = time.Month(monthNow)
	location, err := time.LoadLocation(conf.TimeZoneLocation)
	if err != nil {
		return errors.Wrap(err, "can't parse time zone please check time zone form environment variable")
	}

	startTime := time.Date(yearNow, monthNow, dayNow, 0, 0, 0, 0, location).Unix()
	endTime := time.Date(yearNow, monthNow, dayNow, 0, 0, 0, 0, location).Add(24 * time.Hour).Unix()

	listVisits, err := m.VisitRepo.FindVisitsByDay(startTime, endTime)
	if err != nil {
		return err
	}
	for _, visit := range listVisits {

	}
}
