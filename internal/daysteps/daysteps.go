package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	strings.TrimSpace(datastring)
	data := strings.Split(datastring, ",")

	if len(data) != 2 {
		return fmt.Errorf("Wrong data quantity in dataset")
	}

	for i, p := range data {
		switch i {
		case 0:
			ds.Steps, err = strconv.Atoi(p)
			if err != nil {
				return err
			}
			if ds.Steps <= 0 {
				return fmt.Errorf("steps <= 0")
			}
		case 1:
			ds.Duration, err = time.ParseDuration(p)
			if err != nil {
				return err
			}
			if ds.Duration <= 0 {
				return fmt.Errorf("duration <= 0")
			}
		}

	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	if dist == 0 {
		return "", fmt.Errorf("distance = 0")
	}
	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {

		return "", err

	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, ccal), nil
}
