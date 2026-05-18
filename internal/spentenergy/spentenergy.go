package spentenergy

import (
	"errors"
	"time"
)

var ErrInvalidInput = errors.New("invalid spentenergy input")

const (
	mInKm                 = 1000.0
	minInH                = 60.0
	stepLengthCoefficient = 0.45 // тесты ожидают 0.45
	walkingCoef           = 0.50 // ходьба = 50% от бега
)

// Distance — дистанция (км) по шагам и росту (м).
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	return (float64(steps) * height * stepLengthCoefficient) / mInKm
}

// MeanSpeed — средняя скорость (км/ч).
func MeanSpeed(steps int, height float64, d time.Duration) float64 {
	if d <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	hours := d.Hours()
	if hours <= 0 {
		return 0
	}
	return dist / hours
}

// RunningSpentCalories — калории при беге.
func RunningSpentCalories(steps int, weight, height float64, d time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || d <= 0 {
		return 0, ErrInvalidInput
	}
	ms := MeanSpeed(steps, height, d)
	minutes := d.Minutes()
	cal := (weight * ms * minutes) / minInH
	if cal < 0 {
		cal = 0
	}
	return cal, nil
}

// WalkingSpentCalories — калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, d time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || d <= 0 {
		return 0, ErrInvalidInput
	}
	ms := MeanSpeed(steps, height, d)
	minutes := d.Minutes()
	cal := (weight * ms * minutes) / minInH
	cal *= walkingCoef
	if cal < 0 {
		cal = 0
	}
	return cal, nil
}
