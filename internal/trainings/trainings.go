package trainings

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Training описывает тренировку.
type Training struct {
	Steps        int
	Duration     time.Duration
	TrainingType string
	Personal     personaldata.Personal
}

var ErrInvalidTraining = errors.New("invalid training input")

// Parse парсит строку вида "1000,30m".
func (t *Training) Parse(s string) error {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return ErrInvalidTraining
	}

	var steps int
	if _, err := fmt.Sscanf(parts[0], "%d", &steps); err != nil {
		return err
	}
	if steps <= 0 {
		return ErrInvalidTraining
	}

	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if dur <= 0 {
		return ErrInvalidTraining
	}

	t.Steps = steps
	t.Duration = dur
	return nil
}

// ActionInfo возвращает строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Duration <= 0 || t.Personal.Height <= 0 || t.Personal.Weight <= 0 {
		return "", ErrInvalidTraining
	}

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	// Добавляем объявление переменной cals
	cals, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Шагов: %d\nДлительность: %v\nСкорость: %.2f км/ч\nДистанция: %.2f км\nКалории: %.2f ккал\n",
		t.Steps, t.Duration, speed, dist, cals,
	), nil
}
