package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы

	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, dataString := range dataset {
		err := dp.Parse(dataString)
		if err != nil {
			log.Printf("Ошибка при парсинге данных: %s. Данные: %s\n", err, dataString)
			continue
		}
		infoString, err := dp.ActionInfo()

		if err != nil {
			log.Printf("Ошибка при формировании информации: %s\n", err)
			continue
		}

		log.Println(infoString)
	}
}
