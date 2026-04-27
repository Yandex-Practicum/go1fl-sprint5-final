package personaldata

import "fmt"

type Personal struct {
	Name   string
	weight float64
	height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n")
}
