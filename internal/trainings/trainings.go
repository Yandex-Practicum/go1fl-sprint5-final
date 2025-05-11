package trainings

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go1fl-sprint5-final/internal/personaldata"
	"go1fl-sprint5-final/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		log.Println("Ошибка: нехватка данных")
		return fmt.Errorf("нехватка данных")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		log.Println("Ошибка при преобразовании количества шагов:", err)
		return fmt.Errorf("неверный формат количества шагов")
	}

	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println("Ошибка при преобразовании продолжительности:", err)
		return fmt.Errorf("неверный формат продолжительности")
	}

	if duration <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	if averageSpeed < 0 {
		return "", fmt.Errorf("недопустимая средняя скорость")
	}

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка при расчете калорий: %v", err)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, averageSpeed, calories), nil

}
