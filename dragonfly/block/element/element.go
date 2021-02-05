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
	return Element{element(1), 267}
}
func Helium() Element {
	return Element{element(2), 268}
}
func Lithium() Element {
	return Element{element(3), 269}
}
func Beryllium() Element {
	return Element{element(4), 270}
}
func Boron() Element {
	return Element{element(5), 271}
}
func Carbon() Element {
	return Element{element(6), 272}
}
func Nitrogen() Element {
	return Element{element(7), 273}
}
func Oxygen() Element {
	return Element{element(8), 274}
}
func Fluorine() Element {
	return Element{element(9), 275}
}
func Neon() Element {
	return Element{element(10), 276}
}
func Sodium() Element {
	return Element{element(11), 277}
}
func Magnesium() Element {
	return Element{element(12), 278}
}
func Aluminum() Element {
	return Element{element(13), 279}
}
func Silicon() Element {
	return Element{element(14), 280}
}
func Phosphorus() Element {
	return Element{element(15), 281}
}
func Sulfur() Element {
	return Element{element(16), 282}
}
func Chlorine() Element {
	return Element{element(17), 283}
}
func Argon() Element {
	return Element{element(18), 284}
}
func Potassium() Element {
	return Element{element(19), 285}
}
func Calcium() Element {
	return Element{element(20), 286}
}
func Scandium() Element {
	return Element{element(21), 287}
}
func Titanium() Element {
	return Element{element(22), 288}
}
func Vanadium() Element {
	return Element{element(23), 289}
}
func Chromium() Element {
	return Element{element(24), 290}
}
func Manganese() Element {
	return Element{element(25), 291}
}
func Iron() Element {
	return Element{element(26), 292}
}
func Cobalt() Element {
	return Element{element(27), 293}
}
func Nickel() Element {
	return Element{element(28), 294}
}
func Copper() Element {
	return Element{element(29), 295}
}
func Zinc() Element {
	return Element{element(30), 296}
}
func Gallium() Element {
	return Element{element(31), 297}
}
func Germanium() Element {
	return Element{element(32), 298}
}
func Arsenic() Element {
	return Element{element(33), 299}
}
func Selenium() Element {
	return Element{element(34), 300}
}
func Bromine() Element {
	return Element{element(35), 301}
}
func Krypton() Element {
	return Element{element(36), 302}
}
func Rubidium() Element {
	return Element{element(37), 303}
}
func Strontium() Element {
	return Element{element(38), 304}
}
func Yttrium() Element {
	return Element{element(39), 305}
}
func Zirconium() Element {
	return Element{element(40), 306}
}
func Niobium() Element {
	return Element{element(41), 307}
}
func Molybdenum() Element {
	return Element{element(42), 308}
}
func Technetium() Element {
	return Element{element(43), 309}
}
func Ruthenium() Element {
	return Element{element(44), 310}
}
func Rhodium() Element {
	return Element{element(45), 311}
}
func Palladium() Element {
	return Element{element(46), 312}
}
func Silver() Element {
	return Element{element(47), 313}
}
func Cadmium() Element {
	return Element{element(48), 314}
}
func Indium() Element {
	return Element{element(49), 315}
}
func Tin() Element {
	return Element{element(50), 316}
}
func Antimony() Element {
	return Element{element(51), 317}
}
func Tellurium() Element {
	return Element{element(52), 318}
}
func Iodine() Element {
	return Element{element(53), 319}
}
func Xenon() Element {
	return Element{element(54), 320}
}
func Cesium() Element {
	return Element{element(55), 321}
}
func Barium() Element {
	return Element{element(56), 322}
}
func Lanthanum() Element {
	return Element{element(57), 323}
}
func Cerium() Element {
	return Element{element(58), 324}
}
func Praseodymium() Element {
	return Element{element(59), 325}
}
func Neodymium() Element {
	return Element{element(60), 326}
}
func Promethium() Element {
	return Element{element(61), 327}
}
func Samarium() Element {
	return Element{element(62), 328}
}
func Europium() Element {
	return Element{element(63), 329}
}
func Gadolinium() Element {
	return Element{element(64), 330}
}
func Terbium() Element {
	return Element{element(65), 331}
}
func Dysprosium() Element {
	return Element{element(66), 332}
}
func Holmium() Element {
	return Element{element(67), 333}
}
func Erbium() Element {
	return Element{element(68), 334}
}
func Thulium() Element {
	return Element{element(69), 335}
}
func Ytterbium() Element {
	return Element{element(70), 336}
}
func Lutetium() Element {
	return Element{element(71), 337}
}
func Hafnium() Element {
	return Element{element(72), 338}
}
func Tantalum() Element {
	return Element{element(73), 339}
}
func Tungsten() Element {
	return Element{element(74), 340}
}
func Rhenium() Element {
	return Element{element(75), 341}
}
func Osmium() Element {
	return Element{element(76), 342}
}
func Iridium() Element {
	return Element{element(77), 343}
}
func Platinum() Element {
	return Element{element(78), 344}
}
func Gold() Element {
	return Element{element(79), 345}
}
func Mercury() Element {
	return Element{element(80), 346}
}
func Thallium() Element {
	return Element{element(81), 347}
}
func Lead() Element {
	return Element{element(82), 348}
}
func Bismuth() Element {
	return Element{element(83), 349}
}
func Polonium() Element {
	return Element{element(84), 350}
}
func Astatine() Element {
	return Element{element(85), 351}
}
func Radon() Element {
	return Element{element(86), 352}
}
func Francium() Element {
	return Element{element(87), 353}
}
func Radium() Element {
	return Element{element(88), 354}
}
func Actinium() Element {
	return Element{element(89), 355}
}
func Thorium() Element {
	return Element{element(90), 356}
}
func Protactinium() Element {
	return Element{element(91), 357}
}
func Uranium() Element {
	return Element{element(92), 358}
}
func Neptunium() Element {
	return Element{element(93), 359}
}
func Plutonium() Element {
	return Element{element(94), 360}
}
func Americium() Element {
	return Element{element(95), 361}
}
func Curium() Element {
	return Element{element(96), 362}
}
func Berkelium() Element {
	return Element{element(97), 363}
}
func Californium() Element {
	return Element{element(98), 364}
}
func Einsteinium() Element {
	return Element{element(99), 365}
}
func Fermium() Element {
	return Element{element(100), 366}
}
func Mendelevium() Element {
	return Element{element(101), 367}
}
func Nobelium() Element {
	return Element{element(102), 368}
}
func Lawrencium() Element {
	return Element{element(103), 369}
}
func Rutherfordium() Element {
	return Element{element(104), 370}
}
func Dubnium() Element {
	return Element{element(105), 371}
}
func Seaborgium() Element {
	return Element{element(106), 372}
}
func Bohrium() Element {
	return Element{element(107), 373}
}
func Hassium() Element {
	return Element{element(108), 374}
}
func Meitnerium() Element {
	return Element{element(109), 375}
}
func Darmstadtium() Element {
	return Element{element(110), 376}
}
func Roentgenium() Element {
	return Element{element(111), 377}
}
func Copernicium() Element {
	return Element{element(112), 378}
}
func Nihonium() Element {
	return Element{element(113), 379}
}
func Flerovium() Element {
	return Element{element(114), 380}
}
func Moscovium() Element {
	return Element{element(115), 381}
}
func Livermorium() Element {
	return Element{element(116), 382}
}
func Tennessine() Element {
	return Element{element(117), 383}
}
func Oganesson() Element {
	return Element{element(118), 384}
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
	case "unknown":
		return Element{element(0), 36}, nil
	case "hydrogen":
		return Element{element(1), 267}, nil
	case "helium":
		return Element{element(2), 268}, nil
	case "lithium":
		return Element{element(3), 269}, nil
	case "beryllium":
		return Element{element(4), 270}, nil
	case "boron":
		return Element{element(5), 271}, nil
	case "carbon":
		return Element{element(6), 272}, nil
	case "nitrogen":
		return Element{element(7), 273}, nil
	case "oxygen":
		return Element{element(8), 274}, nil
	case "fluorine":
		return Element{element(9), 275}, nil
	case "neon":
		return Element{element(10), 276}, nil
	case "sodium":
		return Element{element(11), 277}, nil
	case "magnesium":
		return Element{element(12), 278}, nil
	case "aluminum":
		return Element{element(13), 279}, nil
	case "silicon":
		return Element{element(14), 280}, nil
	case "phosphorus":
		return Element{element(15), 281}, nil
	case "sulfur":
		return Element{element(16), 282}, nil
	case "chlorine":
		return Element{element(17), 283}, nil
	case "argon":
		return Element{element(18), 284}, nil
	case "potassium":
		return Element{element(19), 285}, nil
	case "calcium":
		return Element{element(20), 286}, nil
	case "scandium":
		return Element{element(21), 287}, nil
	case "titanium":
		return Element{element(22), 288}, nil
	case "vanadium":
		return Element{element(23), 289}, nil
	case "chromium":
		return Element{element(24), 290}, nil
	case "manganese":
		return Element{element(25), 291}, nil
	case "iron":
		return Element{element(26), 292}, nil
	case "cobalt":
		return Element{element(27), 293}, nil
	case "nickel":
		return Element{element(28), 294}, nil
	case "copper":
		return Element{element(29), 295}, nil
	case "zinc":
		return Element{element(30), 296}, nil
	case "gallium":
		return Element{element(31), 297}, nil
	case "germanium":
		return Element{element(32), 298}, nil
	case "arsenic":
		return Element{element(33), 299}, nil
	case "selenium":
		return Element{element(34), 300}, nil
	case "bromine":
		return Element{element(35), 301}, nil
	case "krypton":
		return Element{element(36), 302}, nil
	case "rubidium":
		return Element{element(37), 303}, nil
	case "strontium":
		return Element{element(38), 304}, nil
	case "yttrium":
		return Element{element(39), 305}, nil
	case "zirconium":
		return Element{element(40), 306}, nil
	case "niobium":
		return Element{element(41), 307}, nil
	case "molybdenum":
		return Element{element(42), 308}, nil
	case "technetium":
		return Element{element(43), 309}, nil
	case "ruthenium":
		return Element{element(44), 310}, nil
	case "rhodium":
		return Element{element(45), 311}, nil
	case "palladium":
		return Element{element(46), 312}, nil
	case "silver":
		return Element{element(47), 313}, nil
	case "cadmium":
		return Element{element(48), 314}, nil
	case "indium":
		return Element{element(49), 315}, nil
	case "tin":
		return Element{element(50), 316}, nil
	case "antimony":
		return Element{element(51), 317}, nil
	case "tellurium":
		return Element{element(52), 318}, nil
	case "iodine":
		return Element{element(53), 319}, nil
	case "xenon":
		return Element{element(54), 320}, nil
	case "cesium":
		return Element{element(55), 321}, nil
	case "barium":
		return Element{element(56), 322}, nil
	case "lanthanum":
		return Element{element(57), 323}, nil
	case "cerium":
		return Element{element(58), 324}, nil
	case "praseodymium":
		return Element{element(59), 325}, nil
	case "neodymium":
		return Element{element(60), 326}, nil
	case "promethium":
		return Element{element(61), 327}, nil
	case "samarium":
		return Element{element(62), 328}, nil
	case "europium":
		return Element{element(63), 329}, nil
	case "gadolinium":
		return Element{element(64), 330}, nil
	case "terbium":
		return Element{element(65), 331}, nil
	case "dysprosium":
		return Element{element(66), 332}, nil
	case "holmium":
		return Element{element(67), 333}, nil
	case "erbium":
		return Element{element(68), 334}, nil
	case "thulium":
		return Element{element(69), 335}, nil
	case "ytterbium":
		return Element{element(70), 336}, nil
	case "lutetium":
		return Element{element(71), 337}, nil
	case "hafnium":
		return Element{element(72), 338}, nil
	case "tantalum":
		return Element{element(73), 339}, nil
	case "tungsten":
		return Element{element(74), 340}, nil
	case "rhenium":
		return Element{element(75), 341}, nil
	case "osmium":
		return Element{element(76), 342}, nil
	case "iridium":
		return Element{element(77), 343}, nil
	case "platinum":
		return Element{element(78), 344}, nil
	case "gold":
		return Element{element(79), 345}, nil
	case "mercury":
		return Element{element(80), 346}, nil
	case "thallium":
		return Element{element(81), 347}, nil
	case "lead":
		return Element{element(82), 348}, nil
	case "bismuth":
		return Element{element(83), 349}, nil
	case "polonium":
		return Element{element(84), 350}, nil
	case "astatine":
		return Element{element(85), 351}, nil
	case "radon":
		return Element{element(86), 352}, nil
	case "francium":
		return Element{element(87), 353}, nil
	case "radium":
		return Element{element(88), 354}, nil
	case "actinium":
		return Element{element(89), 355}, nil
	case "thorium":
		return Element{element(90), 356}, nil
	case "protactinium":
		return Element{element(91), 357}, nil
	case "uranium":
		return Element{element(92), 358}, nil
	case "neptunium":
		return Element{element(93), 359}, nil
	case "plutonium":
		return Element{element(94), 360}, nil
	case "americium":
		return Element{element(95), 361}, nil
	case "curium":
		return Element{element(96), 362}, nil
	case "berkelium":
		return Element{element(97), 363}, nil
	case "californium":
		return Element{element(98), 364}, nil
	case "einsteinium":
		return Element{element(99), 365}, nil
	case "fermium":
		return Element{element(100), 366}, nil
	case "mendelevium":
		return Element{element(101), 367}, nil
	case "nobelium":
		return Element{element(102), 368}, nil
	case "lawrencium":
		return Element{element(103), 369}, nil
	case "rutherfordium":
		return Element{element(104), 370}, nil
	case "dubnium":
		return Element{element(105), 371}, nil
	case "seaborgium":
		return Element{element(106), 372}, nil
	case "bohrium":
		return Element{element(107), 373}, nil
	case "hassium":
		return Element{element(108), 374}, nil
	case "meitnerium":
		return Element{element(109), 375}, nil
	case "darmstadtium":
		return Element{element(110), 376}, nil
	case "roentgenium":
		return Element{element(111), 377}, nil
	case "copernicium":
		return Element{element(112), 378}, nil
	case "nihonium":
		return Element{element(113), 379}, nil
	case "flerovium":
		return Element{element(114), 380}, nil
	case "moscovium":
		return Element{element(115), 381}, nil
	case "livermorium":
		return Element{element(116), 382}, nil
	case "tennessine":
		return Element{element(117), 383}, nil
	case "oganesson":
		return Element{element(118), 384}, nil
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
