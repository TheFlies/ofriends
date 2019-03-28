package types

// Visit hold information of a activity's visit
type Activity struct {
	ID          	string 		`json:"id,omitempty" bson:"_id"`
	StartTime     	string 		`json:"startTime,omitempty"`
	EndTime   		string 		`json:"endTime"`
	Detail  		string    	`json:"detail"`
	Participant    	string     	`json:"participant"`
	Hotel    		string     	`json:"hotel"`
	VisitID			string		`json:"visitId,omitempty" bson:"_visit_id"`
}
