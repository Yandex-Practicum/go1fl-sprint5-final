package spentenergy

import (
	"time"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000.0 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)
//количество калорий для ходьбы
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var err error
	if steps <= 0 {
		return 0, fmt.Errorf("no steps: %w", err)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("no duration: %w", err)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}
//количество калорий для бега
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var err error
	if steps <= 0 {
		return 0, fmt.Errorf("no steps: %w", err)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("no duration: %w", err)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}
//средняя скорость км/ч
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}
//расстояние в км
func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	lengthStep := height * stepLengthCoefficient
	return (float64(steps) * lengthStep) / mInKm
}
