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

	if len(datastring) < 6 {
		return fmt.Errorf("количесто символов меньше ожидаемого")
	}
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("некорректный формат данных: ожидается две части")
	}
	stepsInt, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка подсчета количества шагов")
	}
	if stepsInt <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = stepsInt
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности")
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("неверно задоно время")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("некорректные данные")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
