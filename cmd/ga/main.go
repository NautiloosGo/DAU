package main

import (
	//coll "github.com/NautiloosGo/ga/internal/app/collector"
	"fmt"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	db "github.com/NautiloosGo/ga/internal/services/db"
	"time"
)

// type Dau struct {
// 	Sourse      string
// 	PartnerName string
// 	Date        int
// 	Dau         int
// }
// type Chans struct {
// 	IncomingDau      chan Dau
// 	IncomingPartners chan Dau
// }

var inbox chan ex.Dau

var Channels = ex.Chans{
	IncomingDau:      make(chan ex.Dau),
	IncomingPartners: make(chan ex.Dau),
}

func main() {
	db.FindAndProccessFiles(&Channels)

	for {
		select {
		case NextRow := <-Channels.IncomingDau:
			fmt.Println("received DAU", NextRow)
		case NextRow := <-Channels.IncomingPartners:
			fmt.Println("received Partners", NextRow)
		default:
			time.Sleep(1 * time.Second)
		}
	}

}
