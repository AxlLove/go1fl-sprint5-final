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
	personaldata.Personal

	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	const expParts = 2
	parts := strings.Split(datastring, ",")

	if len(parts) != expParts {
		return fmt.Errorf("invalid input format: expected %d segments, got %d", expParts, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("parse steps: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive, got %d", duration)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("calculate spent calories error: %w", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCal)

	return result, nil
}
