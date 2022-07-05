package collector

import (
	"fmt"
	"time"
)

type Dau struct {
	Sourse      string
	PartnerName string
	Date        int
	Dau         int
}
type Chans struct {
	IncomingDau      chan Dau
	IncomingPartners chan Dau
}

func (c *Chans) Collector() {
	NextRow := ga.Dau{
		Sourse:      "",
		PartnerName: "",
		Date:        0,
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
