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

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",")
	if len(dataParts) != 2 {
		return errors.New("ожидается строка формата шаги,длительность")
	}

	steps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		return fmt.Errorf("некорректная длительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть больше нуля")
	}

	// Записываем данные после всех проверок.
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
