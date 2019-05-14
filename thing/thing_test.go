package thing

import (
	"testing"
	"time"
)

func TestValue(t *testing.T) {
	now = func() time.Time { return time.Unix(1, 0) }
	if got := Value(); got != 1 {
		t.Errorf("Got %v", got)
	}
}
