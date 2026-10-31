package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("the string contains incorrect format")
	}

	steps, error := strconv.Atoi(string(parts[0]))
	if error != nil {
		return fmt.Errorf("conversion failed: %v", error)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	t.Steps = steps
	t.TrainingType = string(parts[1])

	timeDuration, error := time.ParseDuration(string(parts[2]))

	if error != nil {
		return fmt.Errorf("conversion failed: %v", error)
	}

	if timeDuration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	t.Duration = timeDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	calories, err := spentenergy.Calories(t.TrainingType, t.Steps, t.Weight, t.Height, t.Duration)

	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return msg, nil
}
