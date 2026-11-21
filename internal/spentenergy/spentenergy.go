package spentenergy

import (
	"errors"
	"math"
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
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case steps <= 0:
		return 0, errors.New("число шагов должно быть положительным")
	case weight <= 0 || math.IsNaN(weight) || math.IsInf(weight, 0):
		return 0, errors.New("вес должен быть положительным конечным числом")
	case height <= 0 || math.IsNaN(height) || math.IsInf(height, 0):
		return 0, errors.New("рост должен быть положительным конечным числом")
	case duration <= 0:
		return 0, errors.New("длительность должна быть положительной")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	return weight * meanSpeed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
