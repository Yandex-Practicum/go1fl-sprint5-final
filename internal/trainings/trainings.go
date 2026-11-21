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
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	fields := strings.Split(datastring, ",")
	if len(fields) != 3 {
		return errors.New("ожидается запись в формате шаги,тип тренировки,длительность")
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("неверное число шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("число шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(fields[2])
	if err != nil {
		return fmt.Errorf("неверная длительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}

	t.Steps = steps
	t.TrainingType = fields[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
