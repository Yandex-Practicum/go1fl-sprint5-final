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
// Distance рассчитывает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	if steps < 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

// RunningSpentCalories рассчитывает калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	return calories, nil
}

// WalkingSpentCalories рассчитывает калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше нуля")
	}

	if weight <= 0 {
		return 0, errors.New("вес должен быть больше нуля")
	}

	if height <= 0 {
		return 0, errors.New("рост должен быть больше нуля")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше нуля")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}