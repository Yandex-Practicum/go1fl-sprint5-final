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

// DaySteps хранит данные о дневной активности.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m".
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return errors.New("неверный формат данных прогулки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	// Количество шагов должно быть положительным.
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	// Продолжительность должна быть положительной.
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше нуля")
	}

	// Сохраняем данные только после всех проверок.
	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Weight, ds.Height, ds.Duration,
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
