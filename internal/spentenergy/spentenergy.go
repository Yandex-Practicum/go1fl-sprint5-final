package spentenergy

import (
	"errors"
	"time"
)

const (
	stepLengthCoefficient      = 0.45
	mInKm                      = 1000.0
	minInH                     = 60.0
	walkingCaloriesCoefficient = 0.5
)

// Distance возвращает дистанцию в км.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0.0
	}
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0.0
	}
	return distance / hours
}

// RunningSpentCalories рассчитывает калории для бега.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должны быть положительными")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

// WalkingSpentCalories рассчитывает калории для ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("все параметры должны быть положительными")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	baseCalories := (weight * speed * minutes) / minInH
	calories := baseCalories * walkingCaloriesCoefficient
	return calories, nil
}
