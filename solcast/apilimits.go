package solcast

import "time"

type ApiLimits struct {
	ResetTime time.Time
	Limit     int64
	Remaining int64
}
