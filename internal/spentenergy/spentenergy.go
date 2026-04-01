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
	runningSpendCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, fmt.Errorf("err calculate running spend calories: %w", err)
	}
	return runningSpendCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps: %d (must be greater than 0)", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight: %f (must be greater than 0)", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid height: %f (must be greater than 0)", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration: %v (must be positive)", duration)
	}

	durationInMinutes := duration.Minutes()
	meanSpeed := MeanSpeed(steps, height, duration)

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	return (stepLenght * float64(steps)) / mInKm
}
