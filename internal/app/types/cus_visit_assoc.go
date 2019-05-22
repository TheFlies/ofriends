package types

// CusVisitAssoc hold information of a customer visit associate
type CusVisitAssoc struct {
	ID             string `json:"id,omitempty" bson:"_id"`
	CustomerID     string `json:"customerID,omitempty" bson:"_customer_id"`
	VisitID        string `json:"visitID,omitempty" bson:"_visit_id"`
	CustomerName   string `json:"customerName,omitempty"`
	PreApproveVisa bool   `json:"preApproveVisa"`
	CreatedBy      string `json:"createdBy"`
	Note           string `json:"note"`
}
