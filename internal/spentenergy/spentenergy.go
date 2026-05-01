package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка корректности входных параметров
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)
	if speed == 0 {
		return 0, fmt.Errorf("не удалось рассчитать среднюю скорость")
	}

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитываем базовые калории
	baseCalories := (weight * speed * durationInMinutes) / minInH

	// Применяем коэффициент для ходьбы
	calories := baseCalories * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверка корректности входных параметров
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)
	if speed == 0 {
		return 0, fmt.Errorf("не удалось рассчитать среднюю скорость")
	}

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Рассчитываем потраченные калории
	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// Проверка на отрицательные шаги или нулевую/отрицательную продолжительность
	if steps <= 0 || duration <= 0 {
		return 0
	}

	// Вычисляем дистанцию
	distance := Distance(steps, height)

	// Переводим продолжительность в часы
	hours := duration.Hours()

	// Вычисляем среднюю скорость (км/ч)
	speed := distance / hours

	return speed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}
