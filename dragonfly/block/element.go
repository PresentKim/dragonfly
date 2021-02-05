package block

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/block/element"
)

type (
	//Element is a Elements and isotopes are blocks that are essential to chemistry.
	Element struct {
		noNBT
		solid

		//Number is the number of element.
		Number element.Element
	}
)

// EncodeItem ...
func (e Element) EncodeItem() (id int32, meta int16) {
	return e.Number.ItemId, 0
}

// EncodeBlock ...
func (e Element) EncodeBlock() (name string, properties map[string]interface{}) {
	return fmt.Sprintf("minecraft:element_%d", e.Number.Int16()), map[string]interface{}{}
}

// allElements ...
func allElements() (elements []canEncode) {
	for _, e := range element.All() {
		elements = append(elements, Element{Number: e})
	}
	return elements
}
