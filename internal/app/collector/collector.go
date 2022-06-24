package collector

import (
	"fmt"
	"time"
)

type Dau struct {
	Sourse      string
	PartnerName string
	Date        time.Time
	Dau         uint64
}
type Chans struct {
	IncomingDau      chan Dau
	IncomingPartners chan Dau
}

func (c *Chans) Collector() {
	NextRow := ga.Dau{
		Sourse:      "",
		PartnerName: "",
		Date:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Dau:         0,
	}

	for {
		select {
		case NextRow <- c.IncomingDau:
			fmt.Println("received DAU", NextRow.Sourse)
		case NextRow <- c.IncomingDau:
			fmt.Println("received Partners", NextRow.Sourse)
		}
	}
}
