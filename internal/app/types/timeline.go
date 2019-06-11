package types

// Timeline hold information of a visit of customer
type Timeline struct {
	Visit  		Visit
	Customer	*Customer
	Gifts		[] *Gift
	Activities	[] *Activity
}
