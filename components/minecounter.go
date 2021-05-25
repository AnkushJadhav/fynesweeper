package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	xw "fyne.io/x/fyne/widget"
)

const (
	counterSize = 3
)

// MineCounter counts the possible remaining mines
type MineCounter struct {
	*fyne.Container

	digits []*xw.HexWidget
}

// NewMineCounter reates a new mine counter
func NewMineCounter(init int) *MineCounter {
	numd := getDigits(init)
	d := make([]*xw.HexWidget, counterSize)
	c := container.NewHBox()
	for i := 0; i < counterSize; i++ {
		h := xw.NewHexWidget()
		h.SetSlant(0)
		h.Set(numd[i])

		d = append(d, h)
		c.Add(h)
	}

	mc := &MineCounter{
		Container: c,
		digits:    d,
	}
	return mc
}

func getDigits(number int) []uint {
	digits := make([]uint, counterSize)
	for i := 0; i < counterSize; i++ {
		digits[counterSize-i-1] = uint(number % 10)
		number /= 10
	}

	return digits
}
