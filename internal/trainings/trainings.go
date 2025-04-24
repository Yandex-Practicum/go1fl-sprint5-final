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

type calculatedInfo struct {
	distance, meanSpeed, spentCalories float64
}

func (t *Training) Parse(datastring string) (err error) {
	slicedDatastring, err := splitDatastring(datastring)
	if err != nil {
		return err
	}

	t.Steps, err = convertSteps(slicedDatastring[0])
	if err != nil {
		return err
	}

	t.TrainingType = slicedDatastring[1]

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
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	spentCalories, err := t.spentCaloriesByTrainingType()
	if err != nil {
		return "", err
	}

	calcInfo := calculatedInfo{
		distance:      distance,
		meanSpeed:     meanSpeed,
		spentCalories: spentCalories,
	}

	res := t.createInfoResult(calcInfo)

	return res, nil
}

func (t *Training) spentCaloriesByTrainingType() (spentCalories float64, err error) {
	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = errors.New("wrong training type")
	}

	return spentCalories, err
}

func (t *Training) createInfoResult(calcInfo calculatedInfo) string {
	res := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		calcInfo.distance,
		calcInfo.meanSpeed,
		calcInfo.spentCalories)

	return res
}
