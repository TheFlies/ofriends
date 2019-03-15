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
	ArrivedDate     string 		`json:"arrivedDate,omitempty"`
	PreApproveVisa  string 		`json:"preApproveVisa"`
	PassportInfo  	string 		`json:"passportInfo"`
	CreatedBy  		string 		`json:"createdBy"`
	DepartureDate   string 		`json:"departureDate"`
	Activities      []Activity 	`json:"activities"`
}
