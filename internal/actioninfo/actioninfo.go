package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		// Парсим данные
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных: %v\n", err)
			continue
		}

		// Получаем информацию о действии
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации о действии: %v\n", err)
			continue
		}

		// Выводим информацию
		log.Println(info)
	}
}
