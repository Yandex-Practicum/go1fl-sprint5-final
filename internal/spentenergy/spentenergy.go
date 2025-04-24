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

// copy because I don't understand calculations and further objectives about calculations
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateUserParametres(weight, height); err != nil {
		return 0, err
	}

	if err := validateRunningParametres(steps, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := weight * meanSpeed * duration.Minutes() / minInH * walkingCaloriesCoefficient

	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateUserParametres(weight, height); err != nil {
		return 0, err
	}

	if err := validateRunningParametres(steps, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	spentCalories := weight * meanSpeed * duration.Minutes() / minInH

	return spentCalories, nil
}

func validateUserParametres(weight, height float64) error {
	if weight <= 0 {
		return errors.New("weight smaller or equil 0")
	}

	if height <= 0 {
		return errors.New("height smaller or equil 0")
	}

	return nil
}

func validateRunningParametres(steps int, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("steps smaller or equil 0")
	}

	if duration <= 0 {
		return errors.New("duration smaller or equil 0")
	}

	return nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := float64(height * stepLengthCoefficient)
	distance := stepLength * float64(steps) / mInKm
	return distance
}
