package exceller

var keyDAU = "DAU"           //keyword in file name .csv
var keyPartners = "Partners" //keyword in file name .csv

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
type Catalog struct {
	List []Dau
}
