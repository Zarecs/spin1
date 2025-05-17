package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	num, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга количества шагов: %v", err)
	}
	if num <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть положительным")
	}

	trainingTime, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга времени: %v", err)
	}
	if trainingTime <= 0 {
		return 0, 0, fmt.Errorf("время тренировки должно быть положительным")
	}

	return num, trainingTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка парсинга '%s': %v\n", data, err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distanceM := float64(steps) * stepLength
	distanceK := distanceM / mInKm
	caloriesBurned, _ := spentcalories.WalkingSpentCalories(int(distanceK), weight, height, duration)

	result := fmt.Sprintf(`Количество шагов: %d
Дистанция составила: %.2f км
Вы сожгли: %.2f ккал`, steps, distanceK, caloriesBurned)

	return result
}
