package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {

	h := minute / 60

	minute = minute % 60

	if minute < 0 {

		minute += 60

		h--

	}

	hour = (hour + h) % 24

	if hour < 0 {

		hour += 24
	}

	return Clock{hour, minute}

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {

	return New(c.hour, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {

	return New(c.hour, c.minute-minutes)
}
