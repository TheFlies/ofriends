package types

// Gift hold information of a gift
type Gift struct {
	ID          string 		`json:"id,omitempty" bson:"_id"`
	Name        string 		`json:"name,omitempty"`
	Idea        string 		`json:"idea"`
	Size        string 		`json:"size"`
	Quantity    int    		`json:"quantity"`
	Price       float64     `json:"price"`
	Link        string 		`json:"link"`
	Description string 		`json:"description"`
}
