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
	for _, item := range dataset {
		err := dp.Parse(item)
		if err != nil {
			log.Printf("failed to parse %q: %v", item, err)
			continue
		}

		result, err := dp.ActionInfo()
		if err != nil {
			log.Printf("failed to get action info for %q: %v", item, err)
		}

		fmt.Println(result)
	}
}
