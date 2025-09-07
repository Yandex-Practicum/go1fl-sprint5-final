package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	pd "github.com/Yandex-Practicum/tracker/internal/personaldata"
	se "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	pd.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("data size does not match packet size")
	}
	numberOfSteps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("parsing error: %w", err)
	}
	if numberOfSteps <= 0 {
		return fmt.Errorf("wrong number of steps")
	}
	ds.Steps = numberOfSteps
	durations, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("duration parsing error: %w", err)
	}
	if durations <= 0 {
		return fmt.Errorf("wrong duration")
	}
	ds.Duration = durations
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := se.Distance(ds.Steps, ds.Height)
	walkingCallories, err := se.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("walking calorie counting error: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, walkingCallories), nil
}
