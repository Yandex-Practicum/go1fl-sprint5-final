package daysteps

import (
	"fmt"
	"go1fl-sprint5-final/internal/personaldata"
	"go1fl-sprint5-final/internal/spentenergy"
	"log"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		log.Println("Ошибка: нехватка данных")
		return fmt.Errorf("нехватка данных")
	}

	stepsStr := parts[0]
	if strings.TrimSpace(stepsStr) == "" {
		return fmt.Errorf("количество шагов не может быть пустым")
	}

	if strings.HasPrefix(stepsStr, " ") || strings.HasSuffix(stepsStr, " ") {
		return fmt.Errorf("количество шагов не должно содержать пробелов в начале или в конце")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(stepsStr))
	if err != nil {
		log.Println("Ошибка при преобразовании количества шагов:", err)
		return fmt.Errorf("неверный формат количества шагов")
	}

	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	ds.Steps = steps

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println("Ошибка при преобразовании продолжительности:", err)
		return fmt.Errorf("неверный формат продолжительности")
	}

	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше нуля")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)
	return result, nil
}
