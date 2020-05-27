package jumper_transform

import (
	"github.com/jumper86/jumper_transform/impl/transform"
	"github.com/jumper86/jumper_transform/interf"
)

func Newtransform() interf.Transform {
	var tf transform.Transform
	return &tf
}
