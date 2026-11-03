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
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataSplit := strings.Split(datastring, ",")
	if len(dataSplit) != 2 {
		return fmt.Errorf("invalid input : expected 2 fields. Have %d", len(dataSplit))
	}

	stepsInt, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return fmt.Errorf("invalid input %q : expected positive integer. %w", dataSplit[0], err)
	}
	if stepsInt <= 0 {
		return fmt.Errorf("invalid input: steps must be positive. have(%d)", stepsInt)
	}
	ds.Steps = stepsInt

	timeDur, err := time.ParseDuration(dataSplit[1])
	if err != nil {
		return fmt.Errorf("invalid time input: expected 3h00m. %w", err)

	}
	if timeDur <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	ds.Duration = timeDur

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	// TODO: реализовать функцию
	if ds.Steps <= 0 || ds.Duration <= 0 {
		return "", fmt.Errorf("invalid data")
	}
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("impossible to calc calories: %w", err)
	}
	steps := ds.Steps

	res := fmt.Sprintf("Количество шагов: %d.\n"+"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", steps, dist, calories)
	return res, nil

}
