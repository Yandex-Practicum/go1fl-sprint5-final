package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps int
	personaldata.Personal
	Duration time.Duration
}

// создайте метод Parse()
func (d *DaySteps) Parse(smthng string) (err error) {
	split := strings.Split(smthng, ",")

	if len(split) != 2 {
		return errors.New("Длина слайса не равна 2")
	}

	d.Steps, err = strconv.Atoi(split[0])
	if err != nil {
		return errors.New("Ошибка парсинга int")
	}
	d.Duration, err = time.ParseDuration(split[1])
	if err != nil {
		return errors.New("Ошибка парсинга времени")
	}
	return nil
}

// создайте метод ActionInfo()
func (d *DaySteps) ActionInfo() (string, error) {
	if d.Duration <= 0 {
		return "", errors.New("Нет времени. Мыслимо ли это утверждение? Попробуйте помыслить отсутствие времени как такового, прежде чем исправлять данный Error")
	}
	dist := float64(d.Steps) * StepLength

	cal := spentenergy.WalkingSpentCalories(d.Steps, d.Weight, d.Height, d.Duration)

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %2.f км.\nВы сожгли %2.f ккал.\n", d.Steps, dist, cal)

	return info, nil
}
