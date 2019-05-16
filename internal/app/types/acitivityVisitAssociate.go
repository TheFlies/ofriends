package types

// ActivityVisitAssociate hold information of a gift associate
type ActivityVisitAssociate struct {
	ID          string `json:"id,omitempty" bson:"_id"`
	ActivityID  string `json:"activityID,omitempty" bson:"_activity_id"`
	VisitID     string `json:"visitID,omitempty" bson:"_visit_id"`
	Name        string `json:"name,omitempty"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	Participant string `json:"participant"`
	Note        string `json:"note"`
}
