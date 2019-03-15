package service

type (
	Loginservice interface {
		Authenticate(username string, password string) (interface{}, error)
	}
)
