package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal

	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	const expParts = 3
	parts := strings.Split(datastring, ",")

	if len(parts) != expParts {
		return fmt.Errorf("parsing training: expected %d segments, found %d", expParts, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("parse steps: invalid value %q: %w", parts[0], err)
	}

	if steps <= 0 {
		return fmt.Errorf("invalid steps: %d (must be greater than 0)", steps)
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("parse duration: invalid value %s: %w", parts[2], err)
	}

	if duration <= 0 {
		return fmt.Errorf("invalid duration: %v (must be positive)", duration)
	}

	t.Steps = steps
	t.Duration = duration
	t.TrainingType = parts[1]

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var spentCalories float64
	var result string
	var err error

	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return result, fmt.Errorf("error calculate waking spent calories: %w", err)
		}
	case "бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return result, fmt.Errorf("error calculate running spent calories: %w", err)
		}
	default:
		return result, fmt.Errorf("unknown training type")
	}

	result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories)

	return result, nil
}
