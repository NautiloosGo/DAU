package exceller

import (
	"path/filepath"
	"strings"
)

var keyDAU = "DAU"
var keyPartners = "Partners"

type Dau struct {
	Sourse      string
	PartnerName string
	Date        uint64
	Dau         uint64
}
type Chans struct {
	IncomingDau      chan Dau
	IncomingPartners chan Dau
}

func (c *Chans) Switcher(dir string) {
	m, err := filepath.Glob(dir)
	if err != nil {
		panic(err)
	}
	for _, filename := range m {
		if strings.Contains(filename, keyDAU) {
			go c.ReadAndFormatDAU(filename)
		}
		if strings.Contains(filename, keyPartners) {
			go c.ReadAndFormatPartners(filename)
		}
	}

}
