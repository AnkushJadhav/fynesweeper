package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	xw "fyne.io/x/fyne/widget"
)

const (
	timeCounterSize = 3
)

// TimeCounter counts the time
type TimeCounter struct {
	*fyne.Container

	digits  []*xw.HexWidget
	current int
}

// NewTimeCounter creates a new time counter
func NewTimeCounter(init int) *TimeCounter {
	numd := getDigits(init)
	d := make([]*xw.HexWidget, timeCounterSize)
	c := container.NewHBox()
	for i := 0; i < timeCounterSize; i++ {
		h := xw.NewHexWidget()
		h.SetSlant(0)
		h.Set(numd[i])

		d[i] = h
		c.Add(h)
	}

	tc := &TimeCounter{
		Container: c,
		digits:    d,
		current:   init,
	}
	return tc
}

// Increment updates the counter
func (tc *TimeCounter) Increment() {
	tc.current++
	numd := getDigits(tc.current)
	for i := 0; i < timeCounterSize; i++ {
		tc.digits[i].Set(numd[i])
	}
}
