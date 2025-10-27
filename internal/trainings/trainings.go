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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// Разделяем строку по зяпятой
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("должно быть 3 части в строке")
	}

	// Обрабатываем количество шагов
	stepsStr := strings.TrimSpace(parts[0])
	stepsInt, err := strconv.Atoi(stepsStr)
	if err != nil {
		return err
	}
	if stepsInt <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	t.Steps = stepsInt

	// Записываем тип тренировки
	t.TrainingType = strings.TrimSpace(parts[1])

	//Обрабатываем продожительность
	durationStr := strings.TrimSpace(parts[2])
	t.Duration, err = time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	// Проверка на отрицательную или нулевую продолжительность
	if t.Duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	switch t.TrainingType {
	case "Бег":
		calories, _ = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, _ = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
