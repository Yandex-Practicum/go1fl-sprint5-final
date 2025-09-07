package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	pd "github.com/Yandex-Practicum/tracker/internal/personaldata"
	se "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	pd.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("data size does not match packet size")
	}
	numberOfSteps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("parsing error: %w", err)
	}
	if numberOfSteps <= 0 {
		return fmt.Errorf("wrong number of steps")
	}
	t.Steps = numberOfSteps
	t.TrainingType = slice[1]
	durations, err := time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("duration parsing error: %w", err)
	}
	if durations <= 0 {
		return fmt.Errorf("wrong duration")
	}
	t.Duration = durations
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := se.Distance(t.Steps, t.Height)
	averageSpeed := se.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		runningCalories, err := se.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("running calorie calculation error: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, averageSpeed, runningCalories), nil
	case "Ходьба":
		walkingCalories, err := se.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calorie calculations when walking: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), dist, averageSpeed, walkingCalories), nil
	default:
		return "", fmt.Errorf("unknown type of training")
	}
}
