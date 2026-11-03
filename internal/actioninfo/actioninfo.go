package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, activityData := range dataset {
		if err := dp.Parse(activityData); err != nil {
			log.Printf("ошибка разбора данных %q: %v", activityData, err)
			continue
		}

		activityInfo, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка получения информации для %q: %v", activityData, err)
			continue
		}
		fmt.Println(activityInfo)
	}
}
