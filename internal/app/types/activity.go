package types

// Visit hold information of a activity's visit
type Activity struct {
	ID          	string 		`json:"id,omitempty" bson:"_id"`
	StartTime     	int64 		`json:"startTime,omitempty"`
	EndTime   		int64 		`json:"endTime"`
	Detail  		string    	`json:"detail"`
	Participant    	string     	`json:"participant"`
	Hotel    		string     	`json:"hotel"`
	VisitID			string		`json:"visitId,omitempty" bson:"_visit_id"`
}
