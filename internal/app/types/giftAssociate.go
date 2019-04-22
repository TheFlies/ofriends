package types

// GiftAssociate hold information of a gift associate
type GiftAssociate struct {
	ID         string `json:"id,omitempty" bson:"_id"`
	GiftID     string `json:"giftId,omitempty" bson:"_gift_id"`
	FriendID   string `json:"friendId,omitempty" bson:"_friend_id"`
	VisitID    string `json:"visitId,omitempty" bson:"_visit_id"`
	GiftName   string `json:"giftName,omitempty"`
	FriendName string `json:"friendName,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
	Note       string `json:"note"`
}
