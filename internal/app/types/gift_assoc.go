package types

// GiftAssociate hold information of a gift associate
type GiftAssociate struct {
	ID              string `json:"id,omitempty" bson:"_id"`
	GiftID          string `json:"giftID,omitempty" bson:"_gift_id"`
	CusVisitAssocID string `json:"cusvisitassocID,omitempty" bson:"_cusvisitassoc_id"`
	GiftName        string `json:"giftName,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	Note            string `json:"note"`
}
