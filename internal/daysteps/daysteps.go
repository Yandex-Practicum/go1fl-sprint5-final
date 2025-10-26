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
	duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring,",")
	if len(parts) != 2 {
		return errors.New("должно быть две части")
	}
	stepsInt, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if stepsInt <= 0 {
		return err
	}
	ds.Steps = stepsInt
	ds.Duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return errors.New("ошибка парсинга продолжительности:" + err.Error())
	}
	if ds.duration <= 0 {
		return err
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
return fmt.Sprintf("Количество шагов: %d,\nДистанция составила %.2f км.\nВы сохгли %2.f ккал.\n", ds.Steps, distance, calories), nil
}
