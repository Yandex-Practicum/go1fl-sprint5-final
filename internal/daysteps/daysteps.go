package daysteps

import (
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
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("the datastring contains incorrect format")
	}

	steps, error := strconv.Atoi(string(parts[0]))
	if error != nil {
		return fmt.Errorf("conversion failed: %v", error)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	ds.Steps = steps
	timeDuration, error := time.ParseDuration(string(parts[1]))

	if error != nil {
		return fmt.Errorf("conversion failed: %v", error)
	}

	if timeDuration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	ds.Duration = timeDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.Calories(spentenergy.Walk, ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, calories)

	return msg, nil
}
