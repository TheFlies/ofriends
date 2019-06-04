package types

// Timeline hold information of a visit of customer
type Timeline struct {
	Visit  		Visit
	Customers	[] *Customer
	Gifts		[] *Gift
	Activitys	[] *Activity
}
