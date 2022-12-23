package parse2cron

import (
	"fmt"
)

type cronEntity struct {
	cuts [4]int
}

func (ce *cronEntity) Seconds(sec int) *cronEntity {
	ce.cuts[0] += sec
	return ce
}

func (ce *cronEntity) Minutes(min int) *cronEntity {
	ce.cuts[1] += min
	return ce
}

func (ce *cronEntity) Hours(hr int) *cronEntity {
	ce.cuts[2] += hr
	return ce
}

func (ce *cronEntity) End() string {
	for i := 0; i < 3; i++ {
		ce.cuts[i], ce.cuts[i+1] = ce.cuts[i]%60, ce.cuts[i]/60+ce.cuts[i+1]
	}
	return fmt.Sprintf("*/%d */%d */%d * * *", ce.cuts[0], ce.cuts[1], ce.cuts[2])
}

func Seconds(n int) *cronEntity {
	return &cronEntity{
		cuts: [4]int{n, 0, 0},
	}
}

func Minutes(n int) *cronEntity {
	return &cronEntity{
		cuts: [4]int{0, n, 0},
	}
}

func Hours(n int) *cronEntity {
	return &cronEntity{
		cuts: [4]int{0, 0, n},
	}
}

func FromSeconds(n int) string {
	return Seconds(n).End()
}
