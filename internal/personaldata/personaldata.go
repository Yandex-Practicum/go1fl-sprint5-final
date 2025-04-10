package personaldata

import (
	"fmt"
)
type Personal struct {
	Name string //имя пользователя
	Weight float64 //вес пользователя
	Height float64 // рост пользователя
}

func (p Personal) Print() {
	fmt.Sprintf("Имя: %s\nВес: %1.f\nРост: %1.f\n", p.Name, p.Weight, p.Height)
}
