package trainings

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

var ErrInvalidInput = errors.New("invalid training input")

var durationRe = regexp.MustCompile(`^(?:(\d+(?:\.\d+)?)h)?(?:(\d+(?:\.\d+)?)m)?$`)

func parseDurationStrict(s string) (time.Duration, error) {
	if s == "" || strings.Contains(s, " ") {
		return 0, ErrInvalidInput
	}
	m := durationRe.FindStringSubmatch(s)
	if m == nil {
		return 0, ErrInvalidInput
	}
	var hours, mins float64
	var err error
	if m[1] != "" {
		hours, err = strconv.ParseFloat(m[1], 64)
		if err != nil || hours < 0 {
			return 0, ErrInvalidInput
		}
	}
	if m[2] != "" {
		mins, err = strconv.ParseFloat(m[2], 64)
		if err != nil || mins < 0 {
			return 0, ErrInvalidInput
		}
	}
	totalMin := hours*60 + mins
	if totalMin <= 0 {
		return 0, ErrInvalidInput
	}
	return time.Duration(int64(totalMin * 60 * 1e9)), nil
}

func parseStepsStrict(s string) (int, error) {
	if s == "" || strings.Contains(s, " ") {
		return 0, ErrInvalidInput
	}
	if strings.HasPrefix(s, "-") {
		return 0, ErrInvalidInput
	}
	if strings.HasPrefix(s, "+") {
		s = s[1:]
		if s == "" {
			return 0, ErrInvalidInput
		}
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, ErrInvalidInput
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return 0, ErrInvalidInput
	}
	return n, nil
}

func (t *Training) Parse(line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return ErrInvalidInput
	}

	steps, err := parseStepsStrict(parts[0])
	if err != nil {
		return err
	}
	dur, err := parseDurationStrict(parts[2])
	if err != nil {
		return err
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = dur
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if err := t.Personal.Validate(); err != nil {
		return "", err
	}
	switch t.TrainingType {
	case "Бег", "Ходьба":
	default:
		return "", ErrInvalidInput
	}

	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	var cal float64
	var err error
	if t.TrainingType == "Бег" {
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	} else {
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, speed, cal,
	), nil
}

// не обязателен для тестов, но нужен, если main.go вызывает Print()
func (t Training) Print() {
	if s, err := t.ActionInfo(); err == nil {
		fmt.Print(s)
	} else {
		fmt.Println("ошибка:", err)
	}
}
