package personaldata

import "fmt"

// структура с данными пользователя
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Метод ничего не принимает и ничего не возвращает. Он просто выводит данные структуры на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
