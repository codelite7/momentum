package river

import (
	"github.com/riverqueue/river/rivertype"
	"time"
)

type ConstantRetryPolicy struct {
	Delay time.Duration
}

func (p *ConstantRetryPolicy) NextAt(job *rivertype.JobRow) time.Time {
	return time.Now().Add(p.Delay)
}
