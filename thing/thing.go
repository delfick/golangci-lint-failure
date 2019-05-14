package thing

import "time"

var now = time.Now

func Value() int64 {
	return now().Unix()
}
