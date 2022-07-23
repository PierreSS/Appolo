package main

import "time"

func millisTimestampToRFC3339Format(timestamp int64) string {
	return time.Unix(0, timestamp*int64(time.Millisecond)).UTC().Format(time.RFC3339)
}
