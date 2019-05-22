package types

// Visit hold information of a activity's visit
type Activity struct {
	ID     string `json:"id,omitempty" bson:"_id"`
	Name   string `json:"name,omitempty"`
	Detail string `json:"detail"`
}
