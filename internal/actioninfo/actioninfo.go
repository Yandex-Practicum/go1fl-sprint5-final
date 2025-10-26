package actioninfo

import (
	"log"
	"strings"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
ActionInfo() (string,error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for_, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга данных : " + err.Error())
			continue
	}
	info,err := d.ActionInfo()
	if err != nil {
		log.Println("Ошибка при получении информации : " + err.Error())
		continue
	}
	log.Println(info)
	}
}
