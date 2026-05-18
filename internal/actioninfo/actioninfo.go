package actioninfo

import (
	"fmt"
	"log"
)

// интерфейс, который мокается в тесте
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info парсит входной датасет, затем печатает сводную информацию.
func Info(dataset []string, p DataParser) {
	if p == nil || len(dataset) == 0 {
		return
	}
	for _, raw := range dataset {
		if err := p.Parse(raw); err != nil {
			// в тесте happy path ошибок нет; логируем и выходим
			log.Println("parse:", err)
			return
		}
	}
	s, err := p.ActionInfo()
	if err != nil {
		log.Println("action info:", err)
		return
	}
	fmt.Println(s)
}
