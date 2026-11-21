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
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return errors.New("ожидается запись в формате шаги,длительность")
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("неверное число шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("число шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("неверная длительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
