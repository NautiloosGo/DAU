package db

import (
	"github.com/NautiloosGo/ga/internal/app/exceller"
)

// "github.com/NautiloosGo/ga/internal/app/exceller"
const FileDir = "./originaldata/*.csv"

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

func (c *Chans) FindAndProccessFiles() {
	exceller.Switcher(FileDir)
	// exceller.c.Switcher(FileDir)
}
