package types

type Activity struct {
	StartDate     	string 		`json:"startDate,omitempty"`
	EndDate   		string 		`json:"endDate"`
	Detail  		string    	`json:"detail"`
	Participant    	string     	`json:"participant"`
	Hotel    		string     	`json:"hotel"`
}

// Visit hold information of a friend's visit
type Visit struct {
	ID          	string 		`json:"id,omitempty" bson:"_id"`
	Lab        		string 		`json:"lab,omitempty"`
	ArrivedTime     string 		`json:"arrivedTime,omitempty"`
	DepartureTime   string 		`json:"departureTime"`
	PreApproveVisa  bool 		`json:"preApproveVisa"`
	PassportInfo  	string 		`json:"passportInfo"`
	CreatedBy  		string 		`json:"createdBy"`
	HotelStayed  	string 		`json:"hotelStayed"`
	Pickup  		string 		`json:"pickup"`
	Activities      []Activity 	`json:"activities"`
}
