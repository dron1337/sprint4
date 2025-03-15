package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах

)

func parsePackage(data string) (int, time.Duration, error) {
	var duration time.Duration
	dataArr := strings.Split(data, ",")
	if len(dataArr) != 2 {
		return 0, duration, errors.New("не корректные входные данные")
	}
	steps, errAtoi := strconv.Atoi(dataArr[0])
	if errAtoi != nil {
		return 0, duration, errors.New("ошибка преобразования строки в число")
	}
	if steps <= 0 {
		return 0, duration, errors.New("кол-во шагов <= 0")
	}
	duration, errParse := time.ParseDuration(dataArr[1])
	if errParse != nil {
		return 0, duration, errors.New("ошибка преобразования строки в время")
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("%v, произошла в функции: DayActionInfo()\n", err)
		return ""
	}
	distance := float64(steps) * StepLength / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
