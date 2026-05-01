package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат строки")
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}

	if steps <= 0 {
		return errors.New("количество шагов меньше нуля")
	}
	ds.Steps = steps

	// Парсим длительность
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %v", err)
	}

	if duration <= 0 {
		return errors.New("продолжительность меньше нуля")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps <= 0 || ds.Weight <= 0 || ds.Height <= 0 || ds.Duration <= 0 {
		return "", fmt.Errorf("неверные входные параметры: steps, weight, height, и длительность должна быть положительной")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}
	//Сформируйте и верните строку с информацией.
	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)
	return info, nil
}
