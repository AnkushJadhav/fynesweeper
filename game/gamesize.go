package game

// Size stores the game dimensions
type Size struct {
	Rows  int
	Cols  int
	Mines int
}

// Premade difficulty levels
var (
	SizeBeginner = Size{
		Rows:  9,
		Cols:  9,
		Mines: 10,
	}

	SizeIntermediate = Size{
		Rows:  16,
		Cols:  16,
		Mines: 40,
	}

	SizeExpert = Size{
		Rows:  16,
		Cols:  30,
		Mines: 99,
	}
)
