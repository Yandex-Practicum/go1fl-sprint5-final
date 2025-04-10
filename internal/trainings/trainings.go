package trainings

import(
	"strings"
	"time"
	"fmt"
	"strconv"
	"internal/spentenergy"
	"internal/personaldata"
)
type Training struct {
	Steps int //количество шагов
	TrainingType string //тип тренировки
	Duration time.Duration //длительность тренировки
	personaldata.Personal //структура Personal из пакета personaldata
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlise := strings.Split(datastring, ",")
	if len(dataSlise) != 3 {
		return fmt.Errorf("invalid data")
	}
	// выделяем шаги
	steps, err := strconv.Atoi(dataSlise[0])
	if err != nil {
		return fmt.Errorf("conversion steps error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("no steps or error in their quantity: %w", err)
	}
	t.Steps = steps
	t.TrainingType = dataSlise[1]
	duration, err := time.ParseDuration(strings.Replace(dataSlise[3], "h", "h", 1))
	if err != nil {
		return fmt.Errorf("conversion time error: %w", err)
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var distance, meanSpeed, calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		distance = spentenergy.Distance(t.Steps, t.Height)
		meanSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		distance = spentenergy.Distance(t.Steps, t.Height)
		meanSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration), err
	default:
		return "неизвестный тип тренировки", nil
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
