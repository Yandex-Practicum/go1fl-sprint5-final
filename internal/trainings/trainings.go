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
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат данных")
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}

	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// Парсим длительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %v", err)
	}

	if duration <= 0 {
		return errors.New("продолжительность меньше нуля")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// Проверяем валидность данных
	if t.Steps <= 0 {
		return "", fmt.Errorf("некорректное количество шагов: %d", t.Steps)
	}
	if t.Duration <= 0 {
		return "", fmt.Errorf("некорректная длительность: %v", t.Duration)
	}
	if t.Personal.Weight <= 0 || t.Personal.Height <= 0 {
		return "", fmt.Errorf("некорректные персональные данные")
	}

	var calories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчета калорий: %v", err)
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка расчета калорий: %v", err)
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	// Форматируем результат
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
