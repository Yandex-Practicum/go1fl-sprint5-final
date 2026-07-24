package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных: %v", err)
			continue
		}

		str, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка вывода информации об активности: %v", err)
		} else {
			log.Print(str)
		}
		fmt.Println(str)
	}
}
