package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "0", 0, fmt.Errorf("error all data is not specified")
	}

	num, err := strconv.Atoi(strings.TrimSpace(parts[0]))

	if err != nil {
		return 0, "0", 0, fmt.Errorf("error parsing the number of steps: %w", err)
	}

	if num <= 0 {
		return 0, "0", 0, fmt.Errorf("error the number of steps is less than 0")
	}

	trainingTime, err := time.ParseDuration(strings.TrimSpace(parts[2]))

	if trainingTime <= 0 {
		return 0, "0", 0, fmt.Errorf("error negative training time")
	}

	if err != nil {
		return 0, "0", 0, fmt.Errorf("training time parsing error: %w", err)
	}
	typeActivivty := parts[1]
	return num, typeActivivty, trainingTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengStep := height * stepLengthCoefficient
	distM := lengStep * float64(steps)
	distK := distM / mInKm
	return distK
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	durationHours := duration.Hours()
	speedAvg := dist / durationHours
	return speedAvg
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)

	if err != nil {
		return "", fmt.Errorf("error parsin data %w", err)
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64

	if activityType == "Ходьба" {
		calories, _ = WalkingSpentCalories(steps, weight, height, duration)
	} else if activityType == "Бег" {
		calories, _ = RunningSpentCalories(steps, weight, height, duration)
	} else {
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	// Форматируем длительность в часах с двумя знаками после запятой
	durationHours := fmt.Sprintf("%.2f", duration.Hours())

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activityType,
		durationHours,
		dist,
		speed,
		calories,
	), nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) { // ok
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("inсorrect data format for calculating energy. running")
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) { //ok
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect data format for energy calculation. walking")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
