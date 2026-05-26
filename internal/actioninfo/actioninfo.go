package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для парсинга данных и получения информации об активности
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		// Парсим данные
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных в позиции %d ('%s'): %v", i, data, err)
			continue
		}

		// Получаем информацию об активности
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации в позиции %d: %v", i, err)
			continue
		}

		// Выводим информацию
		fmt.Println(info)
	}
}
