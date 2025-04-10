package actioninfo

import (
	"fmt"
	"log"
	"internal/daysteps"
	"internal/trainings"
)
type DataParser interface {
	daysteps.Parse(datastring string) error
	trainings.ActionInfo() string, error
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("invalid data from '%s': %w", data, err)
			continue
		}
		infoAction, err := dp.ActionInfo()
		if err != nil {
			log.Printf("invalid data from '%s': %w", data, err)
			continue
		}

		fmt.Println(infoAction)
	}
}