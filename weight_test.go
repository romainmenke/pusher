package pusher

import (
	"fmt"
	"testing"
	"time"
)

func TestWeight(t *testing.T) {

	fiveMinutesAgo := time.Now().Add(time.Minute * -5)

	p := &push{
		weightedAt: fiveMinutesAgo,
		weight:     10,
	}

	weight(p)

	fmt.Println(p.weight)

}
