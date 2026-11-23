package units

// Inch represents a dimension in inches.
type Inch float64

// Emu represents a dimension in English Metric Units (EMUs).
type Emu int64

type ImageUnits interface {
	ToEmu() Emu
}

// ToEmu converts inches to EMUs.
func (i Inch) ToEmu() Emu {
	return Emu(i * 914400)
}

func (i Emu) ToEmu() Emu {
	return i
}
