package actioninfo

import (
	"fmt"
	"log"
)

// Парсинг
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, entry := range dataset {
		err := dp.Parse(entry)
		if err != nil {
			log.Printf("Ошибка при парсинге данных '%s': %v", entry, err)
			continue
		}

		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка при получении информации об активности для '%s': %v", entry, err)
			continue
		}

		fmt.Println(infoStr)
	}

}
