package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sliceString := strings.Split(datastring, ",")

	if len(sliceString) != 3 {
		return fmt.Errorf("invalid data string: %s", datastring)
	}
	step, err := strconv.Atoi(sliceString[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", sliceString[0])
	}
	ds.Steps = step

	duration, err := time.ParseDuration(sliceString[1])
	if err != nil {
		return fmt.Errorf("invalid duration: %s", sliceString[1])
	}
	ds.Duration = duration

	return err
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	height := float64(ds.Height)

	distance := spentenergy.Distance(ds.Steps, height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories)

	return result, nil
}
