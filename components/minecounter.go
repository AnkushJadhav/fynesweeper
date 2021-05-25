package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	xw "fyne.io/x/fyne/widget"
)

const (
	mineCounterSize = 3
)

// MineCounter counts the possible remaining mines
type MineCounter struct {
	*fyne.Container

	digits  []*xw.HexWidget
	current int
}

// NewMineCounter reates a new mine counter
func NewMineCounter(init int) *MineCounter {
	numd := getDigits(init)
	d := make([]*xw.HexWidget, mineCounterSize)
	c := container.NewHBox()
	for i := 0; i < mineCounterSize; i++ {
		h := xw.NewHexWidget()
		h.SetSize(fyne.NewSize(15, 15))
		h.SetSlant(0)
		h.Set(numd[i])

		d[i] = h
		c.Add(h)
	}

	mc := &MineCounter{
		Container: c,
		digits:    d,
		current:   init,
	}
	return mc
}

// Decrement updates the counter
func (mc *MineCounter) Decrement() {
	mc.current--
	if mc.current < 0 {
		mc.current = 0
	}
	numd := getDigits(mc.current)
	for i := 0; i < mineCounterSize; i++ {
		mc.digits[i].Set(numd[i])
	}
}

func getDigits(number int) []uint {
	digits := make([]uint, mineCounterSize)
	for i := 0; i < mineCounterSize; i++ {
		digits[mineCounterSize-i-1] = uint(number % 10)
		number /= 10
	}

	return digits
}
