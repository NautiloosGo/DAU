package exceller

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func (c *Chans) ReadAndFormatDAU(filename string) {
	NextRow := Dau{
		Sourse:      "DAU",
		PartnerName: "",
		Date:        0,
		Dau:         0,
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)

	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'
	reader.Read()
	reader.Read()

	defer file.Close()

	rname := []rune{}
	for {
		record2, err := reader.Read()
		if err == io.EOF {

			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		rname = []rune{}
		for _, a := range record2[5] {
			if unicode.IsNumber(a) {
				rname = append(rname, a)
			}
		}
		NextRow.Dau, err = strconv.Atoi(string(rname))
		if err != nil {
			fmt.Println(err)
		}

		month, err := strconv.Atoi(string(record2[3]))
		if err != nil {
			fmt.Println(err)
		}
		day, err := strconv.Atoi(string(record2[2]))
		if err != nil {
			fmt.Println(err)
		}
		year, err := strconv.Atoi(string(record2[4]))
		if err != nil {
			fmt.Println(err)
		}
		NextRow.Date = day + month*100 + year*10000
		c.IncomingDau <- NextRow
	}

}
