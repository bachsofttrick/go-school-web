package mytime

import "time"

func GetUnixTime() int64 {
	return time.Now().Unix()
}
