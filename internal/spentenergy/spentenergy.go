package spentenergy

import (
	"fmt"
	"math"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || height <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 ||
		math.IsNaN(weight) || math.IsInf(weight, 0) ||
		math.IsNaN(height) || math.IsInf(height, 0) {
		return 0, fmt.Errorf("некорректные параметры для расчёта калорий при беге")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * meanSpeed * durationInMinutes / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}
