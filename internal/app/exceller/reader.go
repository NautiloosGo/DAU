package exceller

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (c *Chans) ReadAndFormatPartners(fileName string) {

	NextRow := ga.Dau{
		Sourse:      "Partners",
		PartnerName: "",
		Date:        0,
		Dau:         0,
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#' // lines that start with this will be ignored
	reader.Read()        // use Read to remove the first line
	reader.Read()        // use Read to remove the first line

	defer file.Close()

	for {

		record, err := reader.Read()

		if err == io.EOF {
			close(c.chout)
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		rname := []rune{}
		//проверка на вход в каталог или на страницу конкретного партнера (адрес типа /partners/ или /partners )
		// в норме должно быть на входе /partners/'partnerName'  еще могут быть метки или адрес ко: /partners/
		if record[0] == "/partners/" || record[0] == "/partners" || strings.Count(record[0], "/partners?") != 0 {
			NextRow.PartnerName = "Catalogue"
		} else {
			//забираем из строки все что находится на месте %s
			_, err := fmt.Sscanf(record[0], "/partners/%s", &record[0])
			if err != nil {
				//если ошибка, то странная url. Пишем, ERROR  и/или считаем что посетитель только каталога
				record[0] = "ERROR" //точно так нормально делать?
				fmt.Println(err)
			}
			// ищем в url название партнера. Сохраняем все до первого "/"
			for _, a := range record[0] {
				if a == '/' {
					break
				} else {
					rname = append(rname, a)
				}
			}
			// сохраняем название партнера в соответствующую строку в структуре
			NextRow.PartnerName = string(rname)
		}

		massDate := [3]int{}
		massDate[0], err = strconv.Atoi(record[3])
		if err != nil {
			fmt.Println(err)
		}
		massDate[1], err = strconv.Atoi(record[1])
		if err != nil {
			fmt.Println(err)
		}
		massDate[2], err = strconv.Atoi(record[2])
		if err != nil {
			fmt.Println(err)
		}
		NextRow.Date = massDate[0] + massDate[1]*100 + massDate[2]*10000
		NextRow.Dau, err = strconv.Atoi(record[4])
		if err != nil {
			fmt.Println(err)
		}
		// структуру соответствующую одной строке записи передаю в канал. Канал создал отдельно от структуры
		c.inbox <- NextRow
	}
}
