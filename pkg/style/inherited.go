package style

import "github.com/negrel/paon/pkg/pdk/styles/property"

var _ property.Property = InheritedProp{}

type InheritedProp struct {
	property.Prop
}

func Inherited(id property.ID) InheritedProp {
	return InheritedProp{
		Prop: property.MakeProp(id),
	}
}
