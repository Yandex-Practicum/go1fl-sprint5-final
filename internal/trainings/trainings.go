package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return errors.New("invalid steps")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return errors.New("invalid duration")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	if t.Steps <= 0 || t.Duration <= 0 ||
		t.Personal.Weight <= 0 || t.Personal.Height <= 0 {
		return "", errors.New("invalid data")
	}
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
		return "", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil

}
