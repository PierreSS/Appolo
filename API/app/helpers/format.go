package helpers

import "time"

func MillisTimestampToRFC3339Format(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond)).UTC()
}
