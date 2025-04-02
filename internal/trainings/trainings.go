package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(s string) (err error) {
	slice := strings.Split(s, ",")
	if len(slice) != 3 {
		return errors.New("Длина слайса не равна 3")
	}

	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return errors.New("Ошибка парсинга целочисленного значения")
	}

	if slice[1] == "Бег" || slice[1] == "Ходьба" {
		t.TrainingType = slice[1]
	} else {
		return errors.New("Неизвестный вид тренировки")
	}

	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return errors.New("Ошибка парсинга времени")
	}
	return nil
}

// создайте метод ActionInfo()
func (t *Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps)
	var cal float64
	switch t.TrainingType {
	case "Бег":
		cal = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	case "Ходьба":
		cal = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "", errors.New("Неизвестный вид тренировки")
	}

	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	hours := float64(t.Duration) / float64(time.Hour)

	exodus := fmt.Sprintf("Тип тренировки: %s\nДлительность: %2.f ч.\nДистанция: %2.f км.Скорость: %2.f км/ч\nСожгли калорий: %2.f\n", t.TrainingType, hours, dist, speed, cal)
	return exodus, nil
}
