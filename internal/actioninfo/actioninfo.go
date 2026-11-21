package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает типы, которые умеют парсить данные и выводить информацию.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает слайс строк данных.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("ошибка парсинга данных:", err)

			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации:", err)

			continue
		}

		fmt.Println(info)
	}
}