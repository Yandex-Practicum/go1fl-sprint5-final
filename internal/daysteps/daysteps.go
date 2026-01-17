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

// Parse ожидает строку формата: "<steps>,<duration>"
// Примеры: "678,0h50m", "1000,1h30m", "1000,30m", "1000,1.5h".
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	stepsStr := parts[0]
	// По тестам пробелы в начале/конце недопустимы.
	if strings.TrimSpace(stepsStr) != stepsStr {
		return errors.New("invalid steps format")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return errors.New("invalid steps value")
	}

	durStr := parts[1]
	// По тестам пробелы недопустимы (например "1 h30m" должно падать).
	if strings.ContainsAny(durStr, " \t\n\r") {
		return errors.New("invalid duration format")
	}
	dur, err := time.ParseDuration(durStr)
	if err != nil || dur <= 0 {
		return errors.New("invalid duration value")
	}

	ds.Steps = steps
	ds.Duration = dur
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("steps must be positive")
	}
	if ds.Duration <= 0 {
		return "", errors.New("duration must be positive")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight must be positive")
	}
	if ds.Height <= 0 {
		return "", errors.New("height must be positive")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Для daysteps в тестах ожидаются ккал как "ходьба" (коэффициент 0.5),
	// т.е. формула из spentenergy.WalkingSpentCalories.
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
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
