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
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	if datastring == "" {
		return fmt.Errorf("Empty data")
	}
	strings.TrimSpace(datastring)
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("Wrong quantity of fields in dataset")
	}
	for v, field := range data {
		switch v {
		case 0:

			t.Steps, err = strconv.Atoi(field)
			if t.Steps <= 0 {
				return fmt.Errorf("steps <= 0")
			}
			if err != nil {
				return err
			}
		case 1:
			t.TrainingType = field
			if field == "" {
				return fmt.Errorf("неизвестный тип тренировки")
			}
		case 2:
			tim, err := time.ParseDuration(field)
			if err != nil {
				return err
			}
			t.Duration = tim
			if t.Duration <= 0 {
				return fmt.Errorf("duration <= 0")
			}
		}
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var ccal float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		ccal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		ccal, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("Wrong actyvity type")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, speed, ccal), nil
}
