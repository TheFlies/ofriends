package types

type NotificationData struct {
	Customers   []string
	CompanyName string
	DateCome    string
	DateLeave   string
	Location    []string
	HotTime     string
	Req         string
	Note        string
	CreatorName string
}
type VisitsData struct {
	Data []NotificationData
}
