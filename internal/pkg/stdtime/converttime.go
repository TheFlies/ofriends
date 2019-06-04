package stdtime

import "time"

func JSUnixToGo(JSTime int64) time.Time {
	return time.Unix(JSTime/1000, 0)
}

func GoUnixToJS(GoTime time.Time) int64 {
	return GoTime.Unix() * 1000
}
