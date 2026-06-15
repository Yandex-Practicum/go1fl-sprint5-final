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

// Parse ожидает строку формата: "<steps>,<trainingType>,<duration>"
// Примеры: "3456,Ходьба,3h00m", "678,Бег,5m", "+12345,Ходьба,1h30m".
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format")
	}

	stepsStr := parts[0]
	// По тестам пробелы недопустимы.
	if strings.TrimSpace(stepsStr) != stepsStr {
		return errors.New("invalid steps format")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return errors.New("invalid steps value")
	}

	tt := parts[1]
	if tt == "" {
		return errors.New("invalid training type")
	}

	durStr := parts[2]
	// По тестам пробелы в duration недопустимы (например "1 h30m").
	if strings.ContainsAny(durStr, " \t\n\r") {
		return errors.New("invalid duration format")
	}
	dur, err := time.ParseDuration(durStr)
	if err != nil || dur <= 0 {
		return errors.New("invalid duration value")
	}

	t.Steps = steps
	t.TrainingType = tt
	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Базовая валидация входных данных и персональных параметров
	if t.Steps <= 0 {
		return "", errors.New("steps must be positive")
	}
	if t.Duration <= 0 {
		return "", errors.New("duration must be positive")
	}
	if t.Weight <= 0 {
		return "", errors.New("weight must be positive")
	}
	if t.Height <= 0 {
		return "", errors.New("height must be positive")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
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
		meanSpeed,
		calories,
	), nil
}
