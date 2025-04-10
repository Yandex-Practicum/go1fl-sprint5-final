package daysteps

import (
	"fmt"
	"time"
	"internal/spentenergy"
	"internal/personaldata"
	"strconv"
)
type DaySteps struct {
	Steps int //количество шагов
	Duration time.Duration //длительность прогулки
	personaldata.Personal //структура Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlise := strings.Split(datastring, ",")
	if len(dataSlise) != 2 {
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
	ds.Steps = steps	
	// выделяем время
	duration, err := time.ParseDuration(strings.Replace(dataSlise[1], "h", "h", 1))
	if err != nil {
		return fmt.Errorf("conversion time error: %w", err)
	}
	ds.Duration = duration
	return nil	
}

func (ds DaySteps) ActionInfo() (string, error) {
	var distance float64
	distance = spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return fmt.Sprintf("Количество шагов: %s\nДистанция составила: %.2f км.\nВы сожгли: %.2f", ds.Steps, distance, calories), nil
}
