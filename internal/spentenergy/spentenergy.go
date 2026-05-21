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
	switch {
	case steps <= 0:
		return 0.0, fmt.Errorf("количество шагов должно быть больше нуля")
	case weight <= 0:
		return 0.0, fmt.Errorf("вес должен быть больше нуля")
	case height <= 0:
		return 0.0, fmt.Errorf("рост должен быть больше нуля")
	case duration <= 0:
		return 0.0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	switch {
	case steps <= 0:
		return 0.0, fmt.Errorf("количество шагов должно быть больше нуля")
	case weight <= 0:
		return 0.0, fmt.Errorf("вес должен быть больше нуля")
	case height <= 0:
		return 0.0, fmt.Errorf("рост должен быть больше нуля")
	case duration <= 0:
		return 0.0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dis := Distance(steps, height)

	return dis / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}
