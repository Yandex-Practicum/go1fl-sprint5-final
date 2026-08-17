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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sliceString := strings.Split(datastring, ",")

	if len(sliceString) != 3 {
		return fmt.Errorf("invalid data string: %s", datastring)
	}
	id, err := strconv.Atoi(sliceString[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", sliceString[0])
	}
	if id <= 0 {
		return fmt.Errorf("invalid ID: %d", id)
	}
	t.Steps = id

	myTime, err := time.ParseDuration(sliceString[2])
	if err != nil {
		return fmt.Errorf("invalid time: %s", sliceString[2])
	}
	if myTime <= 0 {
		return fmt.Errorf("invalid time: %s", sliceString[2])
	}
	t.Duration = myTime

	t.TrainingType = sliceString[1]

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}

	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	durationHours := t.Duration.Hours()

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationHours, distance, speed, calories), nil
}
