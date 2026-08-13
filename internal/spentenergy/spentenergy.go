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
	if steps < 1 {
		return 0, errors.New("steps = zero")
	}
	if weight < 0 {
		return 0, errors.New("weight = zero")
	}

	if weight == 0 {
		return 0, errors.New("weight = zero")
	}

	if height < 0 {
		return 0, errors.New("height = zero")
	}

	if height == 0 {
		return 0, errors.New("height = zero")
	}

	if duration <= 0 {
		return 0, errors.New("duration = zero")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	walkingSpentCalories := meanSpeed * weight * duration.Hours() * walkingCaloriesCoefficient

	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, errors.New("steps = zero")
	}
	if weight < 0 {
		return 0, errors.New("weight = zero")
	}

	if weight == 0 {
		return 0, errors.New("weight = zero")
	}

	if height < 0 {
		return 0, errors.New("height = zero")
	}

	if height == 0 {
		return 0, errors.New("height = zero")
	}

	if duration <= 0 {
		return 0, errors.New("duration = zero")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	runningSpentCalorie := meanSpeed * weight * duration.Minutes() / minInH

	return runningSpentCalorie, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	meanSpeed := Distance(steps, height) / duration.Hours()

	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps < 0 {
		return 0
	}
	if height < 0 {
		return 0
	}

	steplength := height * stepLengthCoefficient
	distance := float64(steps) * steplength / mInKm

	return distance
}
