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

// Функция возвращает два значения количество калорий, потраченных при ходьбе и ошибку
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Some data <= 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	wsc := ((weight * meanSpeed * duration.Minutes()) / minInH) * walkingCaloriesCoefficient

	return wsc, nil
}

// Func возвращает два значения количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Some data <= 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	rsc := (weight * meanSpeed * duration.Minutes()) / minInH

	return rsc, nil
}

// Функция принимает количество шагов steps, рост пользователя height
// и продолжительность активности duration  и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	speed := Distance(steps, height) / duration.Hours()
	return speed
}

// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLenth := height * stepLengthCoefficient
	distance := stepLenth * float64(steps) / float64(mInKm) // возвращает дистанцию в км
	return distance
}
