package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// проходим по всем строкам
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			// если ошибка парсинга - логируем и идем дальше
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(info)
	}
}
