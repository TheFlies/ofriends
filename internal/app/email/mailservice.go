package email

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/TheFlies/ofriends/internal/app/customer"
	"github.com/TheFlies/ofriends/internal/app/cusvisitassoc"
	"github.com/TheFlies/ofriends/internal/app/types"
	"github.com/TheFlies/ofriends/internal/app/visit"
	envconfig "github.com/TheFlies/ofriends/internal/pkg/config/env"
	"github.com/TheFlies/ofriends/internal/pkg/glog"
	"github.com/TheFlies/ofriends/internal/pkg/mail"
)

type (
	SendMailService struct {
		VisitRepo    visit.Repository
		CustomerRepo customer.Repository
		CusAssVis    cusvisitassoc.Repository
		logger       glog.Logger
	}

	TimeConf struct {
		/*
			the service will send email to notify internal communication the visit is coming (day)
			default is 7days = one week before visit come
		*/
		SendingCondition string `envconfig:"SENDING_CONDITION" default:"7"`
		TimeZoneLocation string `envconfig:"TIME_ZONE_LOCATION" default:"Asia/Bangkok"`
	}
)

const OFRIENDS_DATE_FORMATE = "15:04 Jan 02"

func NewSendMailService(visitRepo visit.Repository, customerRepo customer.Repository, assoc cusvisitassoc.Repository, l glog.Logger) *SendMailService {
	return &SendMailService{
		VisitRepo:    visitRepo,
		CustomerRepo: customerRepo,
		CusAssVis:    assoc,
		logger:       l,
	}
}
func (m *SendMailService) Send(ctx context.Context, now time.Time) error {
	var conf TimeConf
	envconfig.Load(&conf)

	sendingCondition, err := strconv.Atoi(conf.SendingCondition)
	if err != nil {
		m.logger.Errorf("sending condition must a number with integer type")
		return errors.Wrap(err, "the sending condition isn't correct")
	}

	yearNow, monthNow, dayNow := time.Time.Date(now.AddDate(0, 0, sendingCondition))
	monthNow = time.Month(monthNow)
	location, err := time.LoadLocation(conf.TimeZoneLocation)
	if err != nil {
		m.logger.Errorf("can't parse time zone")
		return errors.Wrap(err, "can't parse time zone please check time zone form environment variable")
	}

	startTime := time.Date(yearNow, monthNow, dayNow, 0, 0, 0, 0, location).Unix() * 1000
	endTime := time.Date(yearNow, monthNow, dayNow, 0, 0, 0, 0, location).Add(24*time.Hour).Unix()*1000 - 1

	m.logger.Debugf("start day %d, end day %d", startTime, endTime)
	listVisits, err := m.VisitRepo.FindVisitsByDay(ctx, startTime, endTime)
	if err != nil {
		m.logger.Errorf("%v", err)
		return errors.Wrap(err, "can't send email")
	}
	var totalVisits []types.NotificationData
	for _, visit := range listVisits {
		data := types.NotificationData{}
		data.Location = visit.Lab
		data.DateCome = time.Unix(visit.ArrivedTime/1000, 0).Format(OFRIENDS_DATE_FORMATE)
		data.DateLeave = time.Unix(visit.DepartureTime/1000, 0).Format(OFRIENDS_DATE_FORMATE)
		m.logger.Debugf("time come %v", data.DateCome)
		m.logger.Debugf("time leave %v", data.DateLeave)
		data.CreatorName = "Ofriends domain" // default we should add creator to visit information
		listAssoc, err := m.CusAssVis.FindByVisitID(ctx, visit.ID)
		m.logger.Debugf("list acc :%v", listAssoc)
		if err != nil {
			m.logger.Errorf("%v", err)
			return errors.Wrap(err, "can't send email")
		}
		data.Note = "currently note for visit isn't exits"
		listCustomerName := []string{}
		for _, assoc := range listAssoc {
			customer, err := m.CustomerRepo.FindByID(ctx, assoc.CustomerID)
			if err != nil {
				m.logger.Errorf("%v", err)
				return errors.Wrap(err, "can't send email")
			}
			listCustomerName = append(listCustomerName, customer.Name)
			data.CompanyName = customer.Company

		}
		data.Customers = listCustomerName
		data.Req = "(Interval time to show on TV)"
		data.HotTime = "From 8:30 to 10:30 in the morning From 13:30 to 15:30 in the afternoon"
		totalVisits = append(totalVisits, data)

	}
	m.logger.Debugf("data send mail %v ", totalVisits)
	err = mail.SendEmail(totalVisits)
	if err != nil {
		m.logger.Errorf("%v", err)
		return errors.Wrap(err, "can't send email")
	}
	return nil
}
func (m *SendMailService) SendEmailByVisitID(ctx context.Context, visitID string) error {
	visit, err := m.VisitRepo.FindByID(ctx, visitID)
	if err != nil {
		return errors.Wrap(err, "finding visit in database have error")
	}
	cusAssoc, err := m.CusAssVis.FindByVisitID(ctx, visit.ID)
	customerNames := []string{}
	var companyName string
	for _, c := range cusAssoc {
		customer, err := m.CustomerRepo.FindByID(ctx, c.CustomerID)
		if err != nil {
			return errors.Wrap(err, "getting customer information have error")
		}
		customerNames = append(customerNames, customer.Name)
		companyName = customer.Company
	}

	var listVisits []types.NotificationData
	notificationData := types.NotificationData{
		Customers:   customerNames,
		DateCome:    time.Unix(visit.ArrivedTime/1000, 0).Format(OFRIENDS_DATE_FORMATE),
		DateLeave:   time.Unix(visit.DepartureTime/1000, 0).Format(OFRIENDS_DATE_FORMATE),
		Req:         "(Interval time to show on TV)",
		HotTime:     "From 8:30 to 10:30 in the morning From 13:30 to 15:30 in the afternoon",
		CompanyName: companyName,
		Location:    visit.Lab,
		Note:        "currently note for visit isn't exits",
		CreatorName: "Ofriends domain", // default we should add creator to visit information
	}
	listVisits = append(listVisits, notificationData)
	err = mail.SendEmail(listVisits)
	if err != nil {
		return errors.Wrap(err, "sending email have error")
	}
	return nil
}
