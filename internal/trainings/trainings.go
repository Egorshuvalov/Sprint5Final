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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataSplit := strings.Split(datastring, ",")
	if len(dataSplit) != 3 {
		return fmt.Errorf("invalid input : expected 3 fields. Have %d", len(dataSplit))
	}

	stepsInt, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return fmt.Errorf("invalid input %q : expected positive integer. %w", dataSplit[0], err)
	}
	if stepsInt <= 0 {
		return fmt.Errorf("invalid input: steps must be positive. have(%d)", stepsInt)
	}
	t.Steps = stepsInt

	timeDur, err := time.ParseDuration(dataSplit[2])
	if err != nil {
		return fmt.Errorf("invalid time input: expected 3h00m. %w", err)
	}
	if timeDur <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	t.Duration = timeDur
	training := (dataSplit[1])

	if training == "" {
		return fmt.Errorf("Empty training type")
	}

	t.TrainingType = training
	return nil

}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	active := t.TrainingType
	var calories float64
	var err error

	if active == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("impossible to count calories: %w", err)
		}
	} else if active == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("impossible to count calories: %w", err)
		}
	} else {
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	res := fmt.Sprintf("Тип тренировки: %s\n"+"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n", active, t.Duration.Hours(), dist, speed, calories)

	return res, nil
}
