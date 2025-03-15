package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.

)

func parseTraining(data string) (int, string, time.Duration, error) {
	var duration time.Duration
	dataArr := strings.Split(data, ",")
	if len(dataArr) != 3 {
		return 0, "", duration, errors.New("не корректные входные данные")
	}
	steps, errAtoi := strconv.Atoi(dataArr[0])
	if errAtoi != nil {
		return 0, "", duration, errors.New("ошибка преобразования строки в число")

	}
	duration, errParse := time.ParseDuration(dataArr[2])
	if errParse != nil {
		return 0, "", duration, errors.New("ошибка преобразования строки в время")
	}
	return steps, dataArr[1], duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return lenStep * float64(steps) / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration < 0 {
		return 0.0
	}
	dist := distance(steps)
	return dist / float64(duration.Hours())
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	steps, activityType, duration, err := parseTraining(data)

	var dist, speed, calories float64
	if err != nil {
		fmt.Printf("Ошибка: %v, произошла в функции: DayActionInfo()\n", err)
		return ""
	}
	switch activityType {
	case "Бег":
		dist = distance(steps)
		speed = meanSpeed(steps, duration)
		calories = RunningSpentCalories(steps, weight, duration)

	case "Ходьба":
		dist = distance(steps)
		speed = meanSpeed(steps, duration)
		calories = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "неизвестный тип тренировки"

	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, duration.Hours(), dist, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * float64(duration.Hours()) * minInH
}
