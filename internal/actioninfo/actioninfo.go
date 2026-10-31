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

	if len(dataset) == 0 {
		fmt.Print("")
	} else {
		for _, item := range dataset {
			err := dp.Parse(string(item))

			if err != nil {
				log.Printf("parse item %q: %v", item, err)
				continue
			}
		}

		msg, err := dp.ActionInfo()

		if err != nil {
			log.Printf("%v", err)
		}

		fmt.Print(msg)
	}
}
