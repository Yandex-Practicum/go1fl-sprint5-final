package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчётов.
const (
	mInKm                      = 1000.0 // Количество метров в километре
	minInH                     = 60.0
	stepLengthCoefficient      = 0.45   // Средняя длина шага в метрах для среднего роста человека
	walkingMET                 = 3.5    // Метаболическая эквивалентная трата (MET) при ходьбе
	runningMET                 = 8.0    // Метаболическая эквивалентная трата (MET) при беге
	kcalPerMETMinutePerKg      = 0.0175 // Количество килокалорий, сжигаемых одной MET за минуту для одного кг тела
	walkingCaloriesCoefficient = 0.5
)

// Distance рассчитывает общее расстояние в километрах.
func Distance(steps int, height float64) float64 {
	if steps < 0 || height <= 0 {
		return 0.0
	}

	stepLength := height * stepLengthCoefficient
	totalDistance := float64(steps) * stepLength / mInKm // Переводим в километры

	return totalDistance
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}

	// Проверка отрицательных шагов: если < 0, возвращаем 0
	if steps < 0 {
		return 0.0
	}

	// Проверка роста: если <= 0, возвращаем 0 (некорректные данные)
	if height <= 0 {
		return 0.0
	}

	// Вычисляем дистанцию с помощью функции Distance()
	distance := Distance(steps, height)

	// Переводим продолжительность в часы
	durationHours := duration.Hours()

	// Если durationHours == 0 (из‑за очень малой длительности), избегаем деления на ноль
	if durationHours == 0 {
		return 0.0
	}

	// Рассчитываем среднюю скорость: дистанция (км) / время (ч)
	meanSpeed := distance / durationHours

	return meanSpeed
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка корректности входных параметров
	if steps <= 0 {
		return 0.0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	// Рассчитываем среднюю скорость с помощью MeanSpeed()
	meanSpeed := MeanSpeed(steps, height, duration)

	// Если скорость равна 0 (например, из‑за очень малой длительности), возвращаем ошибку
	if meanSpeed == 0 {
		return 0.0, fmt.Errorf("calculated speed is zero, cannot calculate calories")
	}

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Константа: количество минут в часе

	// Рассчитываем калории по формуле из спецификации:
	// (weight * meanSpeed * durationInMinutes) / minInH
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка корректности входных параметров
	if steps <= 0 {
		return 0.0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	// Рассчитываем среднюю скорость с помощью MeanSpeed()
	meanSpeed := MeanSpeed(steps, height, duration)

	// Если скорость равна 0 (например, из‑за очень малой длительности), возвращаем ошибку
	if meanSpeed == 0 {
		return 0.0, fmt.Errorf("calculated speed is zero, cannot calculate calories")
	}

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитываем калории по основной формуле:
	// (weight * meanSpeed * durationInMinutes) / minInH
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// Применяем корректирующий коэффициент для ходьбы
	calories *= walkingCaloriesCoefficient

	return calories, nil
}
