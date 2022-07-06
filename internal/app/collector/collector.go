package collector

import (
	"fmt"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	db "github.com/NautiloosGo/ga/internal/services/db"
	"sync"
)

func Collector(catalog *db.Catlist, ch ex.Chans) {
	var wg sync.WaitGroup
	go CollectorDAU(catalog, ch, &wg)
	go CollectorPartners(catalog, ch)
}

func CollectorDAU(catalog *db.Catlist, ch ex.Chans, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
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

// 	NextRow := ex.Dau{
// 		Sourse:      "DAU",
// 		PartnerName: "",
// 		Date:        18,
// 		Dau:         112,
// 	}
// 	Catalogs.DAU = append(Catalogs.DAU, NextRow)
// 	Catalogs.Partners = append(Catalogs.Partners, NextRow)
// 	db.RewriteStorage(&Catalogs)

// 	db.GetCatalogs(&Catalogs)

// 	fmt.Println(Catalogs)

// 	db.FindAndProccessFiles(&Channels)

// 	for {
// 		select {
// 		case NextRow := <-Channels.IncomingDau:
// 			fmt.Println("received DAU", NextRow)
// 		case NextRow := <-Channels.IncomingPartners:
// 			fmt.Println("received Partners", NextRow)
// 		default:
// 			time.Sleep(1 * time.Second)
// 		}
// 	}

// }
