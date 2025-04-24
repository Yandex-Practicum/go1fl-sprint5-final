package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, infoObject := range dataset {
		err := dp.Parse(infoObject)
		if err != nil {
			log.Print(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(info)
	}
}
