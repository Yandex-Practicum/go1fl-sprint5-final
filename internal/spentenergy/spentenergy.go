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

var errorSteps = errors.New("The number of steps must be greater than 0")
var errorHeight = errors.New("The height must be greater than 0")
var errorWeight = errors.New("The weight must be greater than 0")
var errorDuration = errors.New("The time interval must be greater than 0")

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errorSteps
	}
	if weight <= 0 {
		return 0, errorWeight
	}
	if height <= 0 {
		return 0, errorHeight
	}
	if duration <= 0 {
		return 0, errorDuration
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errorSteps
	}
	if weight <= 0 {
		return 0, errorWeight
	}
	if height <= 0 {
		return 0, errorHeight
	}
	if duration <= 0 {
		return 0, errorDuration
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / mInKm
}
