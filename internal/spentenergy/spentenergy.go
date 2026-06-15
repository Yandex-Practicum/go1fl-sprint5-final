package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Distance возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLengthMeters := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLengthMeters
	return distanceMeters / mInKm
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	hours := duration.Hours()
	return Distance(steps, height) / hours
}

// RunningSpentCalories возвращает потраченные калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration) // км/ч
	durationMinutes := duration.Minutes()

	// calories = (meanSpeed * weight * minutes) / 60
	calories := (meanSpeed * weight * durationMinutes) / minInH
	return calories, nil
}

// WalkingSpentCalories возвращает потраченные калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration) // км/ч
	durationMinutes := duration.Minutes()

	// calories = (meanSpeed * weight * minutes) / 60 * 0.5
	calories := (meanSpeed * weight * durationMinutes) / minInH * walkingCaloriesCoefficient
	return calories, nil
}
