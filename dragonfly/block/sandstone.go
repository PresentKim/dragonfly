package block

import (
	"github.com/df-mc/dragonfly/dragonfly/block/sandstone"
	"github.com/df-mc/dragonfly/dragonfly/item"
)

type (
	//Sandstone is a solid block commonly found in deserts and beaches underneath sand. Red sandstone is a related block, associated with red sand.
	Sandstone struct {
		noNBT
		solid

		//Type is the type of sandstone.
		Type sandstone.Sandstone
		//Red specifies if the sandstone is red
		Red bool
	}
)

var sandstoneBreakInfo = BreakInfo{
	Hardness:    0.8,
	Harvestable: pickaxeHarvestable,
	Effective:   pickaxeEffective,
}

// BreakInfo ...
func (s Sandstone) BreakInfo() BreakInfo {
	breakInfo := sandstoneBreakInfo
	if s.Type == sandstone.Smooth() {
		breakInfo.Hardness = 2
	}
	breakInfo.Drops = simpleDrops(item.NewStack(s, 1))
	return breakInfo
}

// EncodeItem ...
func (s Sandstone) EncodeItem() (id int32, meta int16) {
	if s.Red {
		id = 179
	} else {
		id = 24
	}
	return id, int16(s.Type.Uint8())
}
