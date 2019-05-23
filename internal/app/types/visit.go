package types

// Visit hold information of a visit
type Visit struct {
	ID            string   `json:"id,omitempty" bson:"_id"`
	Name          string   `json:"name,omitempty"`
	Lab           []string `json:"lab,omitempty"`
	ArrivedTime   int64    `json:"arrivedTime,omitempty"`
	DepartureTime int64    `json:"departureTime,omitempty"`
	Accommodation string   `json:"accommodation"`
	Pickup        string   `json:"pickup"`
}
