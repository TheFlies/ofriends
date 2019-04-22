package types

// Customer hold information of a customer
type Customer struct {
	ID    		    string  `json:"id,omitempty" bson:"_id"`
	Name        	string  `json:"name,omitempty"`
	Title       	string  `json:"title,omitempty"`
	Position    	string  `json:"position,omitempty"`
	Project     	string  `json:"project,omitempty"`
	Age   	    	uint    `json:"age"`
	Company     	string  `json:"company"`
	Country     	string  `json:"country"`
	City        	string  `json:"city"`
	PreApproveVisa  bool 	`json:"preApproveVisa"`
	PassportInfo  	string 	`json:"passportInfo"`
	FoodNote    	string  `json:"foodNote"`
	FamilyNote      string  `json:"familyNote"`
	NextVisitNote   string  `json:"nextVisitNote"`
}
