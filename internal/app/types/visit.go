package types

// Visit hold information of a customer's visit
type Visit struct {
	ID          	string 		`json:"id,omitempty" bson:"_id"`
	Lab        		string 		`json:"lab,omitempty"`
	ArrivedTime     int64 		`json:"arrivedTime,omitempty"`
	DepartureTime   int64		`json:"departureTime,omitempty"`
	PreApproveVisa  bool 		`json:"preApproveVisa"`
	PassportInfo  	string 		`json:"passportInfo"`
	CreatedBy  		string 		`json:"createdBy"`
	HotelStayed  	string 		`json:"hotelStayed"`
	Pickup  		string 		`json:"pickup"`
	CustomerID		string		`json:"customerId,omitempty" bson:"_customer_id"`
}
