package collector

import (
	"fmt"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	db "github.com/NautiloosGo/ga/internal/services/db"
)

func Collector(catalog *db.Catlist, ch ex.Chans) {
	go CollectorDAU(catalog, ch)
	go CollectorPartners(catalog, ch)
}

func CollectorDAU(catalog *db.Catlist, ch ex.Chans) {
	for NextRow := range ch.IncomingDau {
		fmt.Println("received DAU", NextRow)
		AddNewRowDau(catalog, NextRow)
	}
}

func CollectorPartners(catalog *db.Catlist, ch ex.Chans) {
	for NextRow := range ch.IncomingPartners {
		fmt.Println("received Partners", NextRow)
		AddNewRowPartners(catalog, NextRow)
	}
}

func AddNewRowDau(c *db.Catlist, nextrow ex.Dau) {
	for _, d := range c.DAU {
		if nextrow.Date == d.Date {
			break
		}
		c.DAU = append(c.DAU, nextrow)
	}
}

func AddNewRowPartners(c *db.Catlist, nextrow ex.Dau) {
	for _, d := range c.Partners {
		if nextrow.Date == d.Date {
			break
		}
		c.Partners = append(c.Partners, nextrow)
	}
}
