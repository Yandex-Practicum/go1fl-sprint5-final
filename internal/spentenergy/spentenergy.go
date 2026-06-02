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
		return 0, errors.New("invalid input data")
	}

	speed := MeanSpeed(steps, height, duration)

	durationMin := duration.Minutes()

	calories := (weight * speed * durationMin) / minInH

	calories = calories * walkingCaloriesCoefficient

	return calories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input data")
	}
	speed := MeanSpeed(steps, height, duration)

	durationMin := duration.Minutes()

	calories := (weight * speed * durationMin) / minInH

	return calories, nil
}

// Считает среднюю скорость
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

// Distance считает растояние в километрах
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepLength := height * stepLengthCoefficient

	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm

}
