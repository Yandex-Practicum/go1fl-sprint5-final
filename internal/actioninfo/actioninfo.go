package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i := range dataset {
		err := dp.Parse(dataset[i])
		if err != nil {
			log.Println(err)
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(str)
	}

}
