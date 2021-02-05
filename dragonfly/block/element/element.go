package element

import "fmt"

// Element represents a number of elements.
type Element struct {
	Number element
	ItemId int32
}

// Int16 returns the element as a uint8.
func (e Element) Int16() int16 {
	return int16(e.Number)
}

func Unknown() Element {
	return Element{element(0), 36}
}
func Hydrogen() Element {
	return Element{element(1), -12}
}
func Helium() Element {
	return Element{element(2), -13}
}
func Lithium() Element {
	return Element{element(3), -14}
}
func Beryllium() Element {
	return Element{element(4), -15}
}
func Boron() Element {
	return Element{element(5), -16}
}
func Carbon() Element {
	return Element{element(6), -17}
}
func Nitrogen() Element {
	return Element{element(7), -18}
}
func Oxygen() Element {
	return Element{element(8), -19}
}
func Fluorine() Element {
	return Element{element(9), -20}
}
func Neon() Element {
	return Element{element(10), -21}
}
func Sodium() Element {
	return Element{element(11), -22}
}
func Magnesium() Element {
	return Element{element(12), -23}
}
func Aluminum() Element {
	return Element{element(13), -24}
}
func Silicon() Element {
	return Element{element(14), -25}
}
func Phosphorus() Element {
	return Element{element(15), -26}
}
func Sulfur() Element {
	return Element{element(16), -27}
}
func Chlorine() Element {
	return Element{element(17), -28}
}
func Argon() Element {
	return Element{element(18), -29}
}
func Potassium() Element {
	return Element{element(19), -30}
}
func Calcium() Element {
	return Element{element(20), -31}
}
func Scandium() Element {
	return Element{element(21), -32}
}
func Titanium() Element {
	return Element{element(22), -33}
}
func Vanadium() Element {
	return Element{element(23), -34}
}
func Chromium() Element {
	return Element{element(24), -35}
}
func Manganese() Element {
	return Element{element(25), -36}
}
func Iron() Element {
	return Element{element(26), -37}
}
func Cobalt() Element {
	return Element{element(27), -38}
}
func Nickel() Element {
	return Element{element(28), -39}
}
func Copper() Element {
	return Element{element(29), -40}
}
func Zinc() Element {
	return Element{element(30), -41}
}
func Gallium() Element {
	return Element{element(31), -42}
}
func Germanium() Element {
	return Element{element(32), -43}
}
func Arsenic() Element {
	return Element{element(33), -44}
}
func Selenium() Element {
	return Element{element(34), -45}
}
func Bromine() Element {
	return Element{element(35), -46}
}
func Krypton() Element {
	return Element{element(36), -47}
}
func Rubidium() Element {
	return Element{element(37), -48}
}
func Strontium() Element {
	return Element{element(38), -49}
}
func Yttrium() Element {
	return Element{element(39), -50}
}
func Zirconium() Element {
	return Element{element(40), -51}
}
func Niobium() Element {
	return Element{element(41), -52}
}
func Molybdenum() Element {
	return Element{element(42), -53}
}
func Technetium() Element {
	return Element{element(43), -54}
}
func Ruthenium() Element {
	return Element{element(44), -55}
}
func Rhodium() Element {
	return Element{element(45), -56}
}
func Palladium() Element {
	return Element{element(46), -57}
}
func Silver() Element {
	return Element{element(47), -58}
}
func Cadmium() Element {
	return Element{element(48), -59}
}
func Indium() Element {
	return Element{element(49), -60}
}
func Tin() Element {
	return Element{element(50), -61}
}
func Antimony() Element {
	return Element{element(51), -62}
}
func Tellurium() Element {
	return Element{element(52), -63}
}
func Iodine() Element {
	return Element{element(53), -64}
}
func Xenon() Element {
	return Element{element(54), -65}
}
func Cesium() Element {
	return Element{element(55), -66}
}
func Barium() Element {
	return Element{element(56), -67}
}
func Lanthanum() Element {
	return Element{element(57), -68}
}
func Cerium() Element {
	return Element{element(58), -69}
}
func Praseodymium() Element {
	return Element{element(59), -70}
}
func Neodymium() Element {
	return Element{element(60), -71}
}
func Promethium() Element {
	return Element{element(61), -72}
}
func Samarium() Element {
	return Element{element(62), -73}
}
func Europium() Element {
	return Element{element(63), -74}
}
func Gadolinium() Element {
	return Element{element(64), -75}
}
func Terbium() Element {
	return Element{element(65), -76}
}
func Dysprosium() Element {
	return Element{element(66), -77}
}
func Holmium() Element {
	return Element{element(67), -78}
}
func Erbium() Element {
	return Element{element(68), -79}
}
func Thulium() Element {
	return Element{element(69), -80}
}
func Ytterbium() Element {
	return Element{element(70), -81}
}
func Lutetium() Element {
	return Element{element(71), -82}
}
func Hafnium() Element {
	return Element{element(72), -83}
}
func Tantalum() Element {
	return Element{element(73), -84}
}
func Tungsten() Element {
	return Element{element(74), -85}
}
func Rhenium() Element {
	return Element{element(75), -86}
}
func Osmium() Element {
	return Element{element(76), -87}
}
func Iridium() Element {
	return Element{element(77), -88}
}
func Platinum() Element {
	return Element{element(78), -89}
}
func Gold() Element {
	return Element{element(79), -90}
}
func Mercury() Element {
	return Element{element(80), -91}
}
func Thallium() Element {
	return Element{element(81), -92}
}
func Lead() Element {
	return Element{element(82), -93}
}
func Bismuth() Element {
	return Element{element(83), -94}
}
func Polonium() Element {
	return Element{element(84), -95}
}
func Astatine() Element {
	return Element{element(85), -96}
}
func Radon() Element {
	return Element{element(86), -97}
}
func Francium() Element {
	return Element{element(87), -98}
}
func Radium() Element {
	return Element{element(88), -99}
}
func Actinium() Element {
	return Element{element(89), -100}
}
func Thorium() Element {
	return Element{element(90), -101}
}
func Protactinium() Element {
	return Element{element(91), -102}
}
func Uranium() Element {
	return Element{element(92), -103}
}
func Neptunium() Element {
	return Element{element(93), -104}
}
func Plutonium() Element {
	return Element{element(94), -105}
}
func Americium() Element {
	return Element{element(95), -106}
}
func Curium() Element {
	return Element{element(96), -107}
}
func Berkelium() Element {
	return Element{element(97), -108}
}
func Californium() Element {
	return Element{element(98), -109}
}
func Einsteinium() Element {
	return Element{element(99), -110}
}
func Fermium() Element {
	return Element{element(100), -111}
}
func Mendelevium() Element {
	return Element{element(101), -112}
}
func Nobelium() Element {
	return Element{element(102), -113}
}
func Lawrencium() Element {
	return Element{element(103), -114}
}
func Rutherfordium() Element {
	return Element{element(104), -115}
}
func Dubnium() Element {
	return Element{element(105), -116}
}
func Seaborgium() Element {
	return Element{element(106), -117}
}
func Bohrium() Element {
	return Element{element(107), -118}
}
func Hassium() Element {
	return Element{element(108), -119}
}
func Meitnerium() Element {
	return Element{element(109), -120}
}
func Darmstadtium() Element {
	return Element{element(110), -121}
}
func Roentgenium() Element {
	return Element{element(111), -122}
}
func Copernicium() Element {
	return Element{element(112), -123}
}
func Nihonium() Element {
	return Element{element(113), -124}
}
func Flerovium() Element {
	return Element{element(114), -125}
}
func Moscovium() Element {
	return Element{element(115), -126}
}
func Livermorium() Element {
	return Element{element(116), -127}
}
func Tennessine() Element {
	return Element{element(117), -128}
}
func Oganesson() Element {
	return Element{element(118), -129}
}

// All returns a list of all element numbers
func All() []Element {
	return []Element{Unknown(), Hydrogen(), Helium(), Lithium(), Beryllium(), Boron(), Carbon(), Nitrogen(), Oxygen(), Fluorine(), Neon(), Sodium(), Magnesium(), Aluminum(), Silicon(), Phosphorus(), Sulfur(), Chlorine(), Argon(), Potassium(), Calcium(), Scandium(), Titanium(), Vanadium(), Chromium(), Manganese(), Iron(), Cobalt(), Nickel(), Copper(), Zinc(), Gallium(), Germanium(), Arsenic(), Selenium(), Bromine(), Krypton(), Rubidium(), Strontium(), Yttrium(), Zirconium(), Niobium(), Molybdenum(), Technetium(), Ruthenium(), Rhodium(), Palladium(), Silver(), Cadmium(), Indium(), Tin(), Antimony(), Tellurium(), Iodine(), Xenon(), Cesium(), Barium(), Lanthanum(), Cerium(), Praseodymium(), Neodymium(), Promethium(), Samarium(), Europium(), Gadolinium(), Terbium(), Dysprosium(), Holmium(), Erbium(), Thulium(), Ytterbium(), Lutetium(), Hafnium(), Tantalum(), Tungsten(), Rhenium(), Osmium(), Iridium(), Platinum(), Gold(), Mercury(), Thallium(), Lead(), Bismuth(), Polonium(), Astatine(), Radon(), Francium(), Radium(), Actinium(), Thorium(), Protactinium(), Uranium(), Neptunium(), Plutonium(), Americium(), Curium(), Berkelium(), Californium(), Einsteinium(), Fermium(), Mendelevium(), Nobelium(), Lawrencium(), Rutherfordium(), Dubnium(), Seaborgium(), Bohrium(), Hassium(), Meitnerium(), Darmstadtium(), Roentgenium(), Copernicium(), Nihonium(), Flerovium(), Moscovium(), Livermorium(), Tennessine(), Oganesson()}
}

type element uint8

// Uint8 returns the element as a uint8.
func (e element) Uint8() uint8 {
	return uint8(e)
}

// Name ...
func (e element) Name() string {
	switch e {
	case 0:
		return "???"
	case 1:
		return "Hydrogen"
	case 2:
		return "Helium"
	case 3:
		return "Lithium"
	case 4:
		return "Beryllium"
	case 5:
		return "Boron"
	case 6:
		return "Carbon"
	case 7:
		return "Nitrogen"
	case 8:
		return "Oxygen"
	case 9:
		return "Fluorine"
	case 10:
		return "Neon"
	case 11:
		return "Sodium"
	case 12:
		return "Magnesium"
	case 13:
		return "Aluminum"
	case 14:
		return "Silicon"
	case 15:
		return "Phosphorus"
	case 16:
		return "Sulfur"
	case 17:
		return "Chlorine"
	case 18:
		return "Argon"
	case 19:
		return "Potassium"
	case 20:
		return "Calcium"
	case 21:
		return "Scandium"
	case 22:
		return "Titanium"
	case 23:
		return "Vanadium"
	case 24:
		return "Chromium"
	case 25:
		return "Manganese"
	case 26:
		return "Iron"
	case 27:
		return "Cobalt"
	case 28:
		return "Nickel"
	case 29:
		return "Copper"
	case 30:
		return "Zinc"
	case 31:
		return "Gallium"
	case 32:
		return "Germanium"
	case 33:
		return "Arsenic"
	case 34:
		return "Selenium"
	case 35:
		return "Bromine"
	case 36:
		return "Krypton"
	case 37:
		return "Rubidium"
	case 38:
		return "Strontium"
	case 39:
		return "Yttrium"
	case 40:
		return "Zirconium"
	case 41:
		return "Niobium"
	case 42:
		return "Molybdenum"
	case 43:
		return "Technetium"
	case 44:
		return "Ruthenium"
	case 45:
		return "Rhodium"
	case 46:
		return "Palladium"
	case 47:
		return "Silver"
	case 48:
		return "Cadmium"
	case 49:
		return "Indium"
	case 50:
		return "Tin"
	case 51:
		return "Antimony"
	case 52:
		return "Tellurium"
	case 53:
		return "Iodine"
	case 54:
		return "Xenon"
	case 55:
		return "Cesium"
	case 56:
		return "Barium"
	case 57:
		return "Lanthanum"
	case 58:
		return "Cerium"
	case 59:
		return "Praseodymium"
	case 60:
		return "Neodymium"
	case 61:
		return "Promethium"
	case 62:
		return "Samarium"
	case 63:
		return "Europium"
	case 64:
		return "Gadolinium"
	case 65:
		return "Terbium"
	case 66:
		return "Dysprosium"
	case 67:
		return "Holmium"
	case 68:
		return "Erbium"
	case 69:
		return "Thulium"
	case 70:
		return "Ytterbium"
	case 71:
		return "Lutetium"
	case 72:
		return "Hafnium"
	case 73:
		return "Tantalum"
	case 74:
		return "Tungsten"
	case 75:
		return "Rhenium"
	case 76:
		return "Osmium"
	case 77:
		return "Iridium"
	case 78:
		return "Platinum"
	case 79:
		return "Gold"
	case 80:
		return "Mercury"
	case 81:
		return "Thallium"
	case 82:
		return "Lead"
	case 83:
		return "Bismuth"
	case 84:
		return "Polonium"
	case 85:
		return "Astatine"
	case 86:
		return "Radon"
	case 87:
		return "Francium"
	case 88:
		return "Radium"
	case 89:
		return "Actinium"
	case 90:
		return "Thorium"
	case 91:
		return "Protactinium"
	case 92:
		return "Uranium"
	case 93:
		return "Neptunium"
	case 94:
		return "Plutonium"
	case 95:
		return "Americium"
	case 96:
		return "Curium"
	case 97:
		return "Berkelium"
	case 98:
		return "Californium"
	case 99:
		return "Einsteinium"
	case 100:
		return "Fermium"
	case 101:
		return "Mendelevium"
	case 102:
		return "Nobelium"
	case 103:
		return "Lawrencium"
	case 104:
		return "Rutherfordium"
	case 105:
		return "Dubnium"
	case 106:
		return "Seaborgium"
	case 107:
		return "Bohrium"
	case 108:
		return "Hassium"
	case 109:
		return "Meitnerium"
	case 110:
		return "Darmstadtium"
	case 111:
		return "Roentgenium"
	case 112:
		return "Copernicium"
	case 113:
		return "Nihonium"
	case 114:
		return "Flerovium"
	case 115:
		return "Moscovium"
	case 116:
		return "Livermorium"
	case 117:
		return "Tennessine"
	case 118:
		return "Oganesson"
	}
	panic("unknown element number")
}

// FromString ...
func (e element) FromString(s string) (interface{}, error) {
	switch s {
	case "???":
		return Element{element(0), 36}, nil
	case "hydrogen":
		return Element{element(1), -12}, nil
	case "helium":
		return Element{element(2), -13}, nil
	case "lithium":
		return Element{element(3), -14}, nil
	case "beryllium":
		return Element{element(4), -15}, nil
	case "boron":
		return Element{element(5), -16}, nil
	case "carbon":
		return Element{element(6), -17}, nil
	case "nitrogen":
		return Element{element(7), -18}, nil
	case "oxygen":
		return Element{element(8), -19}, nil
	case "fluorine":
		return Element{element(9), -20}, nil
	case "neon":
		return Element{element(10), -21}, nil
	case "sodium":
		return Element{element(11), -22}, nil
	case "magnesium":
		return Element{element(12), -23}, nil
	case "aluminum":
		return Element{element(13), -24}, nil
	case "silicon":
		return Element{element(14), -25}, nil
	case "phosphorus":
		return Element{element(15), -26}, nil
	case "sulfur":
		return Element{element(16), -27}, nil
	case "chlorine":
		return Element{element(17), -28}, nil
	case "argon":
		return Element{element(18), -29}, nil
	case "potassium":
		return Element{element(19), -30}, nil
	case "calcium":
		return Element{element(20), -31}, nil
	case "scandium":
		return Element{element(21), -32}, nil
	case "titanium":
		return Element{element(22), -33}, nil
	case "vanadium":
		return Element{element(23), -34}, nil
	case "chromium":
		return Element{element(24), -35}, nil
	case "manganese":
		return Element{element(25), -36}, nil
	case "iron":
		return Element{element(26), -37}, nil
	case "cobalt":
		return Element{element(27), -38}, nil
	case "nickel":
		return Element{element(28), -39}, nil
	case "copper":
		return Element{element(29), -40}, nil
	case "zinc":
		return Element{element(30), -41}, nil
	case "gallium":
		return Element{element(31), -42}, nil
	case "germanium":
		return Element{element(32), -43}, nil
	case "arsenic":
		return Element{element(33), -44}, nil
	case "selenium":
		return Element{element(34), -45}, nil
	case "bromine":
		return Element{element(35), -46}, nil
	case "krypton":
		return Element{element(36), -47}, nil
	case "rubidium":
		return Element{element(37), -48}, nil
	case "strontium":
		return Element{element(38), -49}, nil
	case "yttrium":
		return Element{element(39), -50}, nil
	case "zirconium":
		return Element{element(40), -51}, nil
	case "niobium":
		return Element{element(41), -52}, nil
	case "molybdenum":
		return Element{element(42), -53}, nil
	case "technetium":
		return Element{element(43), -54}, nil
	case "ruthenium":
		return Element{element(44), -55}, nil
	case "rhodium":
		return Element{element(45), -56}, nil
	case "palladium":
		return Element{element(46), -57}, nil
	case "silver":
		return Element{element(47), -58}, nil
	case "cadmium":
		return Element{element(48), -59}, nil
	case "indium":
		return Element{element(49), -60}, nil
	case "tin":
		return Element{element(50), -61}, nil
	case "antimony":
		return Element{element(51), -62}, nil
	case "tellurium":
		return Element{element(52), -63}, nil
	case "iodine":
		return Element{element(53), -64}, nil
	case "xenon":
		return Element{element(54), -65}, nil
	case "cesium":
		return Element{element(55), -66}, nil
	case "barium":
		return Element{element(56), -67}, nil
	case "lanthanum":
		return Element{element(57), -68}, nil
	case "cerium":
		return Element{element(58), -69}, nil
	case "praseodymium":
		return Element{element(59), -70}, nil
	case "neodymium":
		return Element{element(60), -71}, nil
	case "promethium":
		return Element{element(61), -72}, nil
	case "samarium":
		return Element{element(62), -73}, nil
	case "europium":
		return Element{element(63), -74}, nil
	case "gadolinium":
		return Element{element(64), -75}, nil
	case "terbium":
		return Element{element(65), -76}, nil
	case "dysprosium":
		return Element{element(66), -77}, nil
	case "holmium":
		return Element{element(67), -78}, nil
	case "erbium":
		return Element{element(68), -79}, nil
	case "thulium":
		return Element{element(69), -80}, nil
	case "ytterbium":
		return Element{element(70), -81}, nil
	case "lutetium":
		return Element{element(71), -82}, nil
	case "hafnium":
		return Element{element(72), -83}, nil
	case "tantalum":
		return Element{element(73), -84}, nil
	case "tungsten":
		return Element{element(74), -85}, nil
	case "rhenium":
		return Element{element(75), -86}, nil
	case "osmium":
		return Element{element(76), -87}, nil
	case "iridium":
		return Element{element(77), -88}, nil
	case "platinum":
		return Element{element(78), -89}, nil
	case "gold":
		return Element{element(79), -90}, nil
	case "mercury":
		return Element{element(80), -91}, nil
	case "thallium":
		return Element{element(81), -92}, nil
	case "lead":
		return Element{element(82), -93}, nil
	case "bismuth":
		return Element{element(83), -94}, nil
	case "polonium":
		return Element{element(84), -95}, nil
	case "astatine":
		return Element{element(85), -96}, nil
	case "radon":
		return Element{element(86), -97}, nil
	case "francium":
		return Element{element(87), -98}, nil
	case "radium":
		return Element{element(88), -99}, nil
	case "actinium":
		return Element{element(89), -100}, nil
	case "thorium":
		return Element{element(90), -101}, nil
	case "protactinium":
		return Element{element(91), -102}, nil
	case "uranium":
		return Element{element(92), -103}, nil
	case "neptunium":
		return Element{element(93), -104}, nil
	case "plutonium":
		return Element{element(94), -105}, nil
	case "americium":
		return Element{element(95), -106}, nil
	case "curium":
		return Element{element(96), -107}, nil
	case "berkelium":
		return Element{element(97), -108}, nil
	case "californium":
		return Element{element(98), -109}, nil
	case "einsteinium":
		return Element{element(99), -110}, nil
	case "fermium":
		return Element{element(100), -111}, nil
	case "mendelevium":
		return Element{element(101), -112}, nil
	case "nobelium":
		return Element{element(102), -113}, nil
	case "lawrencium":
		return Element{element(103), -114}, nil
	case "rutherfordium":
		return Element{element(104), -115}, nil
	case "dubnium":
		return Element{element(105), -116}, nil
	case "seaborgium":
		return Element{element(106), -117}, nil
	case "bohrium":
		return Element{element(107), -118}, nil
	case "hassium":
		return Element{element(108), -119}, nil
	case "meitnerium":
		return Element{element(109), -120}, nil
	case "darmstadtium":
		return Element{element(110), -121}, nil
	case "roentgenium":
		return Element{element(111), -122}, nil
	case "copernicium":
		return Element{element(112), -123}, nil
	case "nihonium":
		return Element{element(113), -124}, nil
	case "flerovium":
		return Element{element(114), -125}, nil
	case "moscovium":
		return Element{element(115), -126}, nil
	case "livermorium":
		return Element{element(116), -127}, nil
	case "tennessine":
		return Element{element(117), -128}, nil
	case "oganesson":
		return Element{element(118), -129}, nil
	}

	return nil, fmt.Errorf("unexpected element name '%v'", s)
}

// String ...
func (e element) String() string {
	switch e {
	case 0:
		return "???"
	case 1:
		return "hydrogen"
	case 2:
		return "helium"
	case 3:
		return "lithium"
	case 4:
		return "beryllium"
	case 5:
		return "boron"
	case 6:
		return "carbon"
	case 7:
		return "nitrogen"
	case 8:
		return "oxygen"
	case 9:
		return "fluorine"
	case 10:
		return "neon"
	case 11:
		return "sodium"
	case 12:
		return "magnesium"
	case 13:
		return "aluminum"
	case 14:
		return "silicon"
	case 15:
		return "phosphorus"
	case 16:
		return "sulfur"
	case 17:
		return "chlorine"
	case 18:
		return "argon"
	case 19:
		return "potassium"
	case 20:
		return "calcium"
	case 21:
		return "scandium"
	case 22:
		return "titanium"
	case 23:
		return "vanadium"
	case 24:
		return "chromium"
	case 25:
		return "manganese"
	case 26:
		return "iron"
	case 27:
		return "cobalt"
	case 28:
		return "nickel"
	case 29:
		return "copper"
	case 30:
		return "zinc"
	case 31:
		return "gallium"
	case 32:
		return "germanium"
	case 33:
		return "arsenic"
	case 34:
		return "selenium"
	case 35:
		return "bromine"
	case 36:
		return "krypton"
	case 37:
		return "rubidium"
	case 38:
		return "strontium"
	case 39:
		return "yttrium"
	case 40:
		return "zirconium"
	case 41:
		return "niobium"
	case 42:
		return "molybdenum"
	case 43:
		return "technetium"
	case 44:
		return "ruthenium"
	case 45:
		return "rhodium"
	case 46:
		return "palladium"
	case 47:
		return "silver"
	case 48:
		return "cadmium"
	case 49:
		return "indium"
	case 50:
		return "tin"
	case 51:
		return "antimony"
	case 52:
		return "tellurium"
	case 53:
		return "iodine"
	case 54:
		return "xenon"
	case 55:
		return "cesium"
	case 56:
		return "barium"
	case 57:
		return "lanthanum"
	case 58:
		return "cerium"
	case 59:
		return "praseodymium"
	case 60:
		return "neodymium"
	case 61:
		return "promethium"
	case 62:
		return "samarium"
	case 63:
		return "europium"
	case 64:
		return "gadolinium"
	case 65:
		return "terbium"
	case 66:
		return "dysprosium"
	case 67:
		return "holmium"
	case 68:
		return "erbium"
	case 69:
		return "thulium"
	case 70:
		return "ytterbium"
	case 71:
		return "lutetium"
	case 72:
		return "hafnium"
	case 73:
		return "tantalum"
	case 74:
		return "tungsten"
	case 75:
		return "rhenium"
	case 76:
		return "osmium"
	case 77:
		return "iridium"
	case 78:
		return "platinum"
	case 79:
		return "gold"
	case 80:
		return "mercury"
	case 81:
		return "thallium"
	case 82:
		return "lead"
	case 83:
		return "bismuth"
	case 84:
		return "polonium"
	case 85:
		return "astatine"
	case 86:
		return "radon"
	case 87:
		return "francium"
	case 88:
		return "radium"
	case 89:
		return "actinium"
	case 90:
		return "thorium"
	case 91:
		return "protactinium"
	case 92:
		return "uranium"
	case 93:
		return "neptunium"
	case 94:
		return "plutonium"
	case 95:
		return "americium"
	case 96:
		return "curium"
	case 97:
		return "berkelium"
	case 98:
		return "californium"
	case 99:
		return "einsteinium"
	case 100:
		return "fermium"
	case 101:
		return "mendelevium"
	case 102:
		return "nobelium"
	case 103:
		return "lawrencium"
	case 104:
		return "rutherfordium"
	case 105:
		return "dubnium"
	case 106:
		return "seaborgium"
	case 107:
		return "bohrium"
	case 108:
		return "hassium"
	case 109:
		return "meitnerium"
	case 110:
		return "darmstadtium"
	case 111:
		return "roentgenium"
	case 112:
		return "copernicium"
	case 113:
		return "nihonium"
	case 114:
		return "flerovium"
	case 115:
		return "moscovium"
	case 116:
		return "livermorium"
	case 117:
		return "tennessine"
	case 118:
		return "oganesson"
	}

	panic("unknown element number")
}
