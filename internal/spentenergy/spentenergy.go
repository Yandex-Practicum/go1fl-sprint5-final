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
	if steps <= 0 {
		return 0, fmt.Errorf("input error")
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("input error")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid time value")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	callories := (duration.Minutes() * weight * averageSpeed) / minInH
	walkingCallories := callories * walkingCaloriesCoefficient
	return walkingCallories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("input error")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid time value")
	}
	if weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("input error")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinut := duration.Minutes()
	callories := (weight * averageSpeed * durationInMinut) / minInH
	return callories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps < 0 {
		return 0
	}
	dist := Distance(steps, height)

	averageSpeed := dist / duration.Hours()
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	if steps < 0 {
		return 0
	}
	strideLength := height * stepLengthCoefficient
	distanceInM := float64(steps) * strideLength
	distanceInKm := distanceInM / float64(mInKm)
	return distanceInKm
}
