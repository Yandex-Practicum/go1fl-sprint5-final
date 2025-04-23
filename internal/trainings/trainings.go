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

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// string format: "3456,Ходьба,3h00m"
	slicedDatastring, err := splitDatastring(datastring)
	if err != nil {
		return err
	}

	t.Steps, err = convertSteps(slicedDatastring[0])
	if err != nil {
		return err
	}

	t.TrainingType = slicedDatastring[1]
	if err != nil {
		return err
	}

	t.Duration, err = convertDuration(slicedDatastring[2])
	if err != nil {
		return err
	}

	return err
}

func splitDatastring(datastring string) ([]string, error) {
	slicedDatastring := strings.Split(datastring, ",")
	if len(slicedDatastring) != 3 {
		return []string{}, errors.New("wrong amount of parametres in data string")
	}

	return slicedDatastring, nil
}

func convertSteps(steps string) (int, error) {
	newSteps, err := strconv.Atoi(steps)

	if newSteps <= 0 {
		return 0, errors.New("steps count smaller or equil 0")
	}

	return newSteps, err
}

func convertDuration(duration string) (time.Duration, error) {
	newDuration, err := time.ParseDuration(duration)

	if newDuration <= 0 {
		return 0, errors.New("duration smaller or equil 0")
	}

	return newDuration, err
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var spentCalories float64
	var err error
	if t.TrainingType == "Бег" {
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	res := ""
	res += fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	res += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	res += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	res += fmt.Sprintf("Скорость: %.2f км/ч\n", meanSpeed)
	res += fmt.Sprintf("Сожгли калорий: %.2f\n", spentCalories)

	return "", nil
}
