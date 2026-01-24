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
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func hasLeadingOrTrailingSpace(s string) bool {
	return strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ")
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("ожидается ровно три поля")
	}

	stepStr := parts[0]
	if hasLeadingOrTrailingSpace(stepStr) {
		return errors.New("некорректное значение шагов")
	}
	steps, err := strconv.Atoi(stepStr)
	if err != nil || steps <= 0 {
		return errors.New("некорректное значение шагов")
	}

	trainingType := parts[1]
	durationStr := parts[2]
	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return errors.New("некорректная длительность")
	}

	t.Steps = steps
	t.TrainingType = trainingType
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Weight <= 0 || t.Height <= 0 || t.Duration <= 0 {
		return "", errors.New("некорректные данные")
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n", // ← \n в конце!
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil
}
