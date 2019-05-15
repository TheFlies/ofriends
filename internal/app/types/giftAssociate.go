package types

// GiftAssociate hold information of a gift associate
type GiftAssociate struct {
	ID           string `json:"id,omitempty" bson:"_id"`
	GiftID       string `json:"giftID,omitempty" bson:"_gift_id"`
	CustomerID   string `json:"customerID,omitempty" bson:"_customer_id"`
	VisitID      string `json:"visitID,omitempty" bson:"_visit_id"`
	GiftName     string `json:"giftName,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	Note         string `json:"note"`
}
