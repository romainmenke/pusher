package adaptive

import (
	"math"
	"time"
)

func weight(amount float64, weightedAt time.Time, now time.Time) float64 {
	return amount * (math.Pow((0.5), (now.Sub(weightedAt).Minutes() / 5)))
}
