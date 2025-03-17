package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dron1337/sprint4/internal/spentcalories"
)

const kilometer = 1000

var (
	StepLength = 0.65 // длина шага в метрах

)

func parsePackage(data string) (int, time.Duration, error) {
	var duration time.Duration
	dataArr := strings.Split(data, ",")
	if len(dataArr) != 2 {
		return 0, duration, errors.New("invalid input data")
	}
	steps, errAtoi := strconv.Atoi(dataArr[0])
	if errAtoi != nil {
		return 0, duration, errors.New("invoke strconv.Atoi for steps")
	}
	duration, errParse := time.ParseDuration(dataArr[1])
	if errParse != nil {
		return 0, duration, errors.New("invoke time.ParseDuration for duration")
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
		fmt.Printf("Error: %v, occurred in DayActionInfo()\n", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * StepLength / kilometer
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
