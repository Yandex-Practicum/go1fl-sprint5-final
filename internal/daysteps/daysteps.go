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
	if datastring == "" {
		return errors.New("empty input")
	}

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid format")
	}

	// steps
	stepsStr := parts[0]
	if strings.HasPrefix(stepsStr, "+") {
		stepsStr = strings.TrimPrefix(stepsStr, "+")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return errors.New("invalid steps")
	}

	// duration
	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return errors.New("invalid duration")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("invalid data")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		calories,
	), nil
}
