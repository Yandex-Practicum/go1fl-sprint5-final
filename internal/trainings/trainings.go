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
	Personal     personaldata.Personal
}

func (t Training) Print() {
	t.Personal.Print()
}

func (t *Training) Parse(datastring string) error {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("Данные введены неверно")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("Неверное количество шагов - отрицательное значение")
	}
	t.Steps = steps

	t.TrainingType = data[1]

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("Неверная продолжительность")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Personal.Weight,
			t.Personal.Height,
			t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Personal.Weight,
			t.Personal.Height,
			t.Duration,
		)
	default:
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	durationHours := t.Duration.Hours()
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		durationHours,
		distance,
		speed,
		calories,
	)
	return result, nil
}
