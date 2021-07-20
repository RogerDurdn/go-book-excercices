package main

import (
	"time"
)

func main() {
	print(daysToFriday() + "\n")
	print(switchNot())
}

func daysToFriday() string {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 1:
		return "tomorrow"
	case today + 2:
		return "in 2 days"
	default:
		return "good luck"

	}
}

func switchNot() (x string) {
	t := time.Now()
	switch {
	case t.Hour() > 12:
		x = "late " + t.Local().String()
	case t.Hour() < 12:
		x = "tempra" + t.Local().String()
	}
	return
}
