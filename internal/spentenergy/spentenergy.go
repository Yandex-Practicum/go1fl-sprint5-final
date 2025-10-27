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

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("входные данные: шаги, вес, рост, время тренировки для расчета калорий указаны неверно")
	}
	average := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentcalories := (weight * average * minutes) / minInH
	return spentcalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("входные данные: шаги, вес, рост, время тренировки для расчета калорий указаны неверно")
	}
	average := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentcalories := (weight * average * minutes) / minInH
	return spentcalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distanceKilometers := Distance(steps, height)
	hours := duration.Hours()
	return distanceKilometers / hours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthStep := height * stepLengthCoefficient
	distanceMeters := float64(steps) * lengthStep
	return distanceMeters / mInKm
}
