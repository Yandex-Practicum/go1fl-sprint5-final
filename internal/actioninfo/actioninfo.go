package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, d DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := d.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга данных : " + err.Error())
			continue
		}
		info, err := d.ActionInfo()
		if err != nil {
			log.Println("ошибка при получении информации : " + err.Error())
			continue
		}
		log.Println(info)
	}
}
