package main

import (
	"fmt"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	db "github.com/NautiloosGo/ga/internal/services/db"
	"time"
)

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
