package main

import (
	"time"

	db "github.com/NautiloosGo/ga/internal/services/db"
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

var inbox chan Dau

var Channels = Chans{
	IncomingDau:      make(chan Dau),
	IncomingPartners: make(chan Dau),
}

func main() {
	db.Channels.FindAndProccessFiles()
}
