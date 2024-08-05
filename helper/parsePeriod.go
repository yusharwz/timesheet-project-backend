package helper

import (
	"errors"
	"strconv"
	"strings"
)

func ParsePeriod(period string) error {
	var err error
	var startPeriod int
	var endPeriod int

	result := strings.Split(period, ":")
	startPeriod, err = strconv.Atoi(result[0])
	if err != nil {
		return err
	} else if startPeriod < 1 || startPeriod > 12 {
		return errors.New("startPeriod must be greater than 0 and less than 12")
	}

	endPeriod, err = strconv.Atoi(result[1])
	if err != nil {
		return err
	} else if endPeriod > 12 || endPeriod < 1 {
		return errors.New("endPeriod must be greater than 0 and less than 12")
	}
	return nil
}
