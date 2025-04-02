package actioninfo

import (
	"fmt"
	"log"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, s := range dataset {
		err := dp.Parse(s)
		if err != nil {
			log.Printf("Ошибка парсинга данных: %v, строка: %s\n", err, s)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка %v, строка: %s\n", err, info)
			continue
		}
		fmt.Println(info)
	}
}
