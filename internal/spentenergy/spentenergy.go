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
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * averageSpeed * duration.Minutes()) / minInH
	calories = calories * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	averageSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * averageSpeed * duration.Minutes()) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)

	averageSpeed := distance / duration.Hours()

	return averageSpeed
}

func Distance(steps int, height float64) float64 {

	strideLength := height * stepLengthCoefficient
	distance := float64(steps) * strideLength
	distance = distance / mInKm
	return distance
}
