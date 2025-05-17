package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	// Обработка количества шагов
	stepsStr := parts[0]
	if stepsStr == "" {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("количество шагов не указано")
	}

	// Проверка на наличие только знака + или -
	if stepsStr == "+" || stepsStr == "-" {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("неверный формат количества шагов")
	}

	// Обработка знака "+"
	if strings.HasPrefix(stepsStr, "+") {
		stepsStr = stepsStr[1:]
	}
	// проверка на пробелы
	trimmedStepsStr := strings.TrimSpace(stepsStr)
	if trimmedStepsStr != stepsStr {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("количество шагов не должно содержать пробелов в начале или конце")
	}

	num, err := strconv.Atoi(stepsStr)
	if err != nil {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("ошибка парсинга количества шагов: %v", err)
	}
	if num <= 0 {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	// Обработка времени
	timeStr := parts[1]
	if timeStr == "" {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("время тренировки не указано")
	}

	// Проверка на пробелы внутри времени
	if strings.Contains(timeStr, " ") {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("неверный формат продолжительности")
	}

	trainingTime, err := time.ParseDuration(timeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга времени: %v", err)
	}
	if trainingTime <= 0 {
		log.Println("сообщение")
		return 0, 0, fmt.Errorf("время тренировки должно быть положительным")
	}

	return num, trainingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	if data == "" {
		log.Println("сообщение")
		return ""
	}

	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceK := distanceM / mInKm
	caloriesBurned, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceK, caloriesBurned)
}
