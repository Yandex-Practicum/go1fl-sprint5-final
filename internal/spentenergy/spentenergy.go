package spentenergy

import (
	"errors"
	"time"
)

const (
	// длина шага = 0.45 * рост (м)
	stepLengthCoefficient = 0.45
	mInKm                 = 1000.0
	minInH                = 60.0
	// коэффициент для ходьбы
	walkingCaloriesCoefficient = 0.5
)

var ErrInvalidInput = errors.New("invalid input parameters")

// Distance — дистанция в километрах.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := stepLengthCoefficient * height // м/шаг
	distMeters := float64(steps) * stepLength    // м
	return distMeters / mInKm                    // км
}

// MeanSpeed — средняя скорость (км/ч).
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	dist := Distance(steps, height) // км
	hours := duration.Hours()       // ч
	if hours <= 0 {
		return 0
	}
	return dist / hours // км/ч
}

// RunningSpentCalories — калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrInvalidInput
	}
	speed := MeanSpeed(steps, height, duration) // км/ч
	minutes := duration.Minutes()               // мин
	cal := (weight * speed * minutes) / minInH
	return cal, nil
}

// WalkingSpentCalories — калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrInvalidInput
	}
	speed := MeanSpeed(steps, height, duration) // км/ч
	minutes := duration.Minutes()               // мин
	base := (weight * speed * minutes) / minInH
	return base * walkingCaloriesCoefficient, nil
}
