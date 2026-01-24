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
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

// hasLeadingOrTrailingSpace проверяет, есть ли пробелы в начале или конце
func hasLeadingOrTrailingSpace(s string) bool {
	return strings.HasPrefix(s, " ") || strings.HasSuffix(s, " ")
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("ожидается ровно два поля")
	}

	stepStr := parts[0]
	if hasLeadingOrTrailingSpace(stepStr) {
		return errors.New("некорректное значение шагов")
	}
	steps, err := strconv.Atoi(stepStr)
	if err != nil || steps <= 0 {
		return errors.New("некорректное значение шагов")
	}

	durationStr := parts[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return errors.New("некорректная длительность")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Weight <= 0 || ds.Height <= 0 || ds.Duration <= 0 {
		return "", errors.New("некорректные данные")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n", // ← здесь добавлен \n в конце!
		ds.Steps,
		distance,
		calories,
	), nil
}
