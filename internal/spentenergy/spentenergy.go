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

const (
	Run  = "Бег"
	Walk = "Ходьба"
)

func checkParameters(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("steps number must be positive")
	}

	if weight <= 0 {
		return fmt.Errorf("weight must be positive")
	}

	if height <= 0 {
		return fmt.Errorf("weight must be positive")
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := checkParameters(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	avgSpeed := MeanSpeed(steps, height, duration)

	return ((weight * avgSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	err := checkParameters(steps, weight, height, duration)

	if err != nil {
		return 0, err
	}

	avgSpeed := MeanSpeed(steps, height, duration)

	return (duration.Minutes() * weight * avgSpeed) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	avgSpeed := distance / duration.Hours()

	return avgSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	distance := (float64(steps) * stepLength) / mInKm

	return distance
}

func Calories(trainingType string, steps int, weight float64, height float64, duration time.Duration) (float64, error) {
	switch trainingType {
	case Run:
		calories, err := RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return 0, err
		}

		return calories, nil

	case Walk:
		calories, err := WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return 0, err
		}

		return calories, nil
	default:
		return 0, fmt.Errorf("неизвестный тип тренировки")
	}
}
