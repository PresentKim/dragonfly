package sandstone

import "fmt"

// SandStone represents a type of  block. Some blocks, such as sandstone, sandstone stairs carry one of these types.
type Sandstone struct {
	sandstone
}

// Default returns default sandstone material.
func Default() Sandstone {
	return Sandstone{sandstone(0)}
}

// Chiseled returns chiseled sandstone material.
func Chiseled() Sandstone {
	return Sandstone{sandstone(1)}
}

// Cut returns cut sandstone material.
func Cut() Sandstone {
	return Sandstone{sandstone(2)}
}

// Smooth returns smooth sandstone material.
func Smooth() Sandstone {
	return Sandstone{sandstone(3)}
}

// All returns a list of all sandstone types
func All() []Sandstone {
	return []Sandstone{Default(), Chiseled(), Cut(), Smooth()}
}

type sandstone uint8

// Uint8 returns the sandstone as a uint8.
func (ss sandstone) Uint8() uint8 {
	return uint8(ss)
}

// Name ...
func (ss sandstone) Name() string {
	switch ss {
	case 0:
		return "Default"
	case 1:
		return "Chiseled"
	case 2:
		return "Cut"
	case 3:
		return "Smooth"
	}
	panic("unknown sandstone type")
}

// FromString ...
func (ss sandstone) FromString(s string) (interface{}, error) {
	switch s {
	case "default":
		return Sandstone{sandstone(0)}, nil
	case "heiroglyphs":
		return Sandstone{sandstone(1)}, nil
	case "cut":
		return Sandstone{sandstone(2)}, nil
	case "smooth":
		return Sandstone{sandstone(3)}, nil
	}
	return nil, fmt.Errorf("unexpected sandstone type '%v', expecting one of 'default', 'heiroglyphs', 'cut', 'smooth'", s)
}

// String ...
func (ss sandstone) String() string {
	switch ss {
	case 0:
		return "default"
	case 1:
		return "chiseled"
	case 2:
		return "cut"
	case 3:
		return "smooth"
	}
	panic("unknown sandstone type")
}
