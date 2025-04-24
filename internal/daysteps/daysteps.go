package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

type calculatedInfo struct {
	distance, spentCalories float64
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitedData, err := splitDatastring(datastring)
	if err != nil {
		return err
	}

	ds.Steps, err = parseSteps(splitedData[0])
	if err != nil {
		return err
	}

	ds.Duration, err = parseDuration(splitedData[1])
	if err != nil {
		return err
	}

	return nil
}

func splitDatastring(datastring string) ([]string, error) {
	splitedData := strings.Split(datastring, ",")
	if len(splitedData) != 2 {
		return []string{}, errors.New("wrong amount of paramtres")
	}

	return splitedData, nil
}

func parseSteps(steps string) (int, error) {
	newSteps, err := strconv.Atoi(steps)
	if newSteps <= 0 {
		return 0, errors.New("steps smaller or equil 0")
	}

	return newSteps, err
}

func parseDuration(duration string) (time.Duration, error) {
	newDuration, err := time.ParseDuration(duration)
	if newDuration <= 0 {
		return 0, errors.New("duration smaller or equil 0")
	}

	return newDuration, err
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	calcInfo := calculatedInfo{
		distance:      distance,
		spentCalories: spentCalories,
	}

	res := ds.createInfoResult(calcInfo)

	return res, nil
}

func (ds DaySteps) createInfoResult(calcInfo calculatedInfo) string {
	res := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		calcInfo.distance,
		calcInfo.spentCalories)

	return res
}
