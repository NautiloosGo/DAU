package main

import (
	"fmt"
	col "github.com/NautiloosGo/ga/internal/app/collector"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	db "github.com/NautiloosGo/ga/internal/services/db"
)

var Channels = ex.Chans{
	IncomingDau:      make(chan ex.Dau),
	IncomingPartners: make(chan ex.Dau),
}

var Catalogs = db.Catlist{
	DAU:      make([]ex.Dau, 0),
	Partners: make([]ex.Dau, 0),
}

func main() {
	//download json
	db.GetCatalogs(&Catalogs)
	fmt.Println(Catalogs)

	//start download data (parsing .csv)
	db.FindAndProccessFiles(&Channels)

	//manage inbox chans
	col.Collector(&Catalogs, Channels)

	// for {
	// 	select {
	// 	case NextRow := <-Channels.IncomingDau:
	// 		fmt.Println("received DAU", NextRow)
	// 		col.AddNewRowDau(&Catalogs, NextRow)
	// 	case NextRow := <-Channels.IncomingPartners:
	// 		fmt.Println("received Partners", NextRow)
	// 		col.AddNewRowPartners(&Catalogs, NextRow)
	// 	default:
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }

	db.RewriteStorage(&Catalogs)

}
