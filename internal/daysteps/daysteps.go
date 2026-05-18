package daysteps

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

type DaySteps struct {
	Steps                 int
	Duration              time.Duration
	personaldata.Personal // встраиваем: поля Weight/Height «поднимаются» наружу
}

var ErrInvalidInput = errors.New("invalid daysteps input")

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

func (d *DaySteps) Parse(line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return ErrInvalidInput
	}
	steps, err := parseStepsStrict(parts[0])
	if err != nil {
		return err
	}
	dur, err := parseDurationStrict(parts[1])
	if err != nil {
		return err
	}
	d.Steps = steps
	d.Duration = dur
	return nil
}

func (d DaySteps) ActionInfo() (string, error) {
	if d.Steps <= 0 || d.Duration <= 0 {
		return "", ErrInvalidInput
	}
	if err := d.Personal.Validate(); err != nil {
		return "", err
	}

	dist := spentenergy.Distance(d.Steps, d.Height)
	cal, err := spentenergy.WalkingSpentCalories(d.Steps, d.Weight, d.Height, d.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		d.Steps, dist, cal,
	), nil
}
