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
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("steps, weight, height, duration must be positive. Your input:\n"+
			"steps = %-10d\n"+"weight = %-10.2f\n"+"height = %-10.2f\n"+"duration = %-10s\n", steps, weight, height, duration.String())
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentCalories := (weight * speed * minutes) / float64(minInH) * walkingCaloriesCoefficient
	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("steps, weight, height, duration must be positive. Your input:\n"+
			"steps = %-10d\n"+"weight = %-10.2f\n"+"height = %-10.2f\n"+"duration = %-10s\n", steps, weight, height, duration.String())
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	spentCalories := (weight * speed * minutes) / float64(minInH)
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	durHrs := duration.Hours()

	dist := Distance(steps, height)
	meanSpeed := dist / durHrs
	return meanSpeed

}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLen := height * stepLengthCoefficient
	distInMeters := float64(steps) * stepLen
	dist := distInMeters / float64(mInKm)
	return dist
}
