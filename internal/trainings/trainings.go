package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training содержит данные о тренировке и информацию о пользователе
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")

	// Проверяем, что у нас ровно 3 части: шаги, тип, длительность
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format: expected 3 fields, got %d", len(parts))
	}

	// Парсим количество шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse steps: %v", err)
	}
	// Валидация шагов: должны быть положительными
	if steps <= 0 {
		return fmt.Errorf("invalid steps count: %d (must be positive)", steps)
	}
	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = parts[1]

	// Парсим длительность тренировки
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("failed to parse duration: %v", err)
	}
	// Валидация длительности: должна быть положительной
	if duration <= 0 {
		return fmt.Errorf("invalid duration: %v (must be positive)", duration)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)

	// Вычисляем среднюю скорость — функция не возвращает ошибку
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	// Рассчитываем калории в зависимости от типа тренировки
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate walking calories: %v", err)
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate running calories: %v", err)
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	// Форматируем длительность в часах (с двумя знаками после запятой)
	durationInHours := t.Duration.Hours()

	// Формируем итоговую строку
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		durationInHours,
		distance,
		meanSpeed,
		calories)

	return result, nil
}
